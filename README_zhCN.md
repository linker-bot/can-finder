# CAN Finder 项目说明

[English](./README.md)

## 项目概述

**CAN Finder** 是一款专为灵心巧手设备设计的网络设备自动发现和管理工具。该项目通过监听设备发送的 UDP 广播，实时展示设备的基本信息，包括设备名、IP 地址、MAC 地址、设备型号和版本信息，并通过 WebSocket 协议向前端实时推送更新。

## 功能特性

* **实时设备发现**：通过监听 UDP 广播数据，实时发现网络中的 CAN 设备。
* **Web 界面展示**：通过友好的 Web 界面直观展示设备详细信息，包含首次发现和最后活跃时间。
* **WebSocket 通信**：利用 WebSocket 协议实时更新设备状态，无需手动刷新。

## 系统架构

CAN Finder 由两个主要部分组成：

1. **后端服务**：采用 Go 语言编写，使用 Gin Web 框架提供 HTTP 服务、WebSocket 服务和 UDP 广播监听。
2. **前端展示页面**：采用原生 HTML 和 JavaScript 构建，通过 WebSocket 与后端实时交互。

## 技术栈

* Go 语言
* Gin 框架
* WebSocket (Gorilla)
* HTML 和 JavaScript
* UDP 广播协议

## 使用方法

### 安装运行

```shell
go build -o can-finder
./can-finder --http-port 6200 --udp-port 9999
```

### 访问 Web 界面

启动服务后，通过浏览器访问：

```
http://localhost:6200
```

即可实时查看 CAN 设备的发现情况。

## 项目结构

```
can-finder/
├── main.go             # 后端核心程序，处理UDP广播及HTTP/WebSocket请求
└── public
    └── index.html      # 前端展示页面
```

## 配置说明

* HTTP 服务默认监听端口：`6200`（可通过命令行参数修改）
* UDP 广播监听默认端口：`9999`（可通过命令行参数修改）

## 示例页面展示

前端展示页面实时更新设备信息，包括：

* 设备名
* IP 地址
* MAC 地址
* 设备型号
* 软件版本（链接至 GitHub Release 页面）
* 首次发现时间
* 最后活跃时间

## 错误处理与日志

后端服务提供丰富的日志信息，以快速排查设备发现及通信相关问题。

## 项目依赖

* github.com/gin-gonic/gin
* github.com/gorilla/websocket
* github.com/soulteary/gin-static

使用 `go mod tidy` 自动管理依赖。

## License

本项目采用 GPL-3.0 license。

---

欢迎贡献代码、提出 Issue 或提交 PR，帮助我们不断优化和完善 CAN Finder！
