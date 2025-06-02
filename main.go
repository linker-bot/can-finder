package main

import (
	"encoding/json"
	"log"
	"net"
)

// DeviceInfo 接收到的设备信息结构体
type DeviceInfo struct {
	Name    string `json:"name"`
	IP      string `json:"ip"`
	MAC     string `json:"mac"`
	Model   string `json:"model"`
	Version string `json:"version"`
}

func main() {
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
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("接收失败: %v", err)
			continue
		}

		var device DeviceInfo
		if err := json.Unmarshal(buf[:n], &device); err != nil {
			log.Printf("JSON解析失败: %v", err)
			continue
		}

		log.Printf("发现设备 %s (%s) 来自 %s, 型号: %s, 版本: %s",
			device.Name, device.IP, remoteAddr.IP, device.Model, device.Version)
	}
}

