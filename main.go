package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	static "github.com/soulteary/gin-static"
)

//go:embed public
var EmbedFS embed.FS

type DeviceInfo struct {
	Name    string `json:"name"`
	IP      string `json:"ip"`
	MAC     string `json:"mac"`
	Model   string `json:"model"`
	Version string `json:"version"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func udpListener() {
	addr := net.UDPAddr{
		Port: 9999,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		log.Fatalf("监听UDP失败: %v", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	log.Println("开始监听设备广播...")

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("接收失败: %v", err)
			continue
		}

		var device DeviceInfo
		if err := json.Unmarshal(buf[:n], &device); err != nil {
			log.Printf("JSON解析失败: %v", err)
			continue
		}

		data, _ := json.Marshal(device)
		for client := range clients {
			client.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("升级失败:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	defer delete(clients, conn)

	for {
		if _, _, err := conn.NextReader(); err != nil {
			break
		}
	}
}

func main() {
	go udpListener()

	r := gin.Default()

	r.GET("/", static.ServeEmbed("public", EmbedFS))

	r.GET("/ws", wsHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/config", func(c *gin.Context) {
		callback := c.Query("callback")
		host := c.Request.Host

		if callback == "" {
			c.JSON(400, gin.H{"error": "callback parameter required"})
			return
		}

		data := gin.H{
			"host": host,
		}

		c.JSONP(200, data)
	})

	r.NoRoute(func(c *gin.Context) {
		fmt.Printf("%s doesn't exists, redirect on /\n", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":18080")
}
