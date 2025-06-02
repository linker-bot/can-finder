package main

import (
	"embed"
	"encoding/json"
	"flag"
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

// DeviceInfo holds device information from UDP broadcasts
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

// udpListener listens for UDP broadcasts from devices
func udpListener(port int) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		log.Fatalf("❌ UDP listener failed to start on port %d: %v", port, err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	log.Printf("📡 Listening for device broadcasts on UDP port %d", port)

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("⚠️ Error receiving UDP broadcast: %v", err)
			continue
		}

		var device DeviceInfo
		if err := json.Unmarshal(buffer[:n], &device); err != nil {
			log.Printf("⚠️ Failed to parse device info: %v", err)
			continue
		}

		data, _ := json.Marshal(device)
		for client := range clients {
			client.WriteMessage(websocket.TextMessage, data)
		}
	}
}

// wsHandler handles WebSocket upgrade requests and client management
func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("⚠️ WebSocket upgrade failed: %v", err)
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

// main initializes and starts the HTTP and UDP services
func main() {
	var httpPort int
	var udpPort int

	flag.IntVar(&httpPort, "http-port", 6200, "Port for HTTP server")
	flag.IntVar(&udpPort, "udp-port", 9999, "Port for UDP listener")
	flag.Parse()

	log.Printf("🚀 Starting device discovery service")

	go udpListener(udpPort)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", static.ServeEmbed("public", EmbedFS))
	r.GET("/ws", wsHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/config", func(c *gin.Context) {
		callback := c.Query("callback")
		host := c.Request.Host

		if callback == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "callback parameter required"})
			return
		}

		data := gin.H{"host": host}
		c.JSONP(http.StatusOK, data)
	})

	r.NoRoute(func(c *gin.Context) {
		log.Printf("🔄 Redirecting unknown route '%s' to '/'", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	log.Printf("🌐 HTTP server running on port %d", httpPort)
	r.Run(fmt.Sprintf(":%d", httpPort))
}
