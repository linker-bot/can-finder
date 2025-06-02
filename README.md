# CAN Finder Project Description

[中文文档](./README_zhCN.md)

## Project Overview

**CAN Finder** is a network device auto-discovery and management tool specifically designed for LinkerHand devices. The tool listens for UDP broadcasts from devices, displays their essential information in real-time, such as device name, IP address, MAC address, model, and version information, and pushes updates to the frontend via the WebSocket protocol.

## Features

* **Real-Time Device Discovery**: Listens to UDP broadcast data to discover CAN devices in real-time.
* **Web Interface Display**: Provides an intuitive web interface to show detailed device information, including first discovery time and last active time.
* **WebSocket Communication**: Uses WebSocket to update device status in real-time without manual refresh.

## System Architecture

CAN Finder consists of two main components:

1. **Backend Service**: Written in Go, using the Gin web framework to provide HTTP services, WebSocket services, and UDP broadcast listening.
2. **Frontend Display Page**: Built with native HTML and JavaScript, communicating with the backend in real-time through WebSockets.

## Technology Stack

* Go Language
* Gin Framework
* WebSocket (Gorilla)
* HTML and JavaScript
* UDP Broadcast Protocol

## Usage

### Installation and Execution

```shell
go build -o can-finder
./can-finder --http-port 6200 --udp-port 9999
```

### Accessing the Web Interface

After starting the service, access it through a browser at:

```
http://localhost:6200
```

This will allow real-time monitoring of CAN device discovery.

## Project Structure

```
can-finder/
├── main.go             # Backend core program, handling UDP broadcasts and HTTP/WebSocket requests
└── public
    └── index.html      # Frontend display page
```

## Configuration Details

* Default HTTP server listening port: `6200` (modifiable via command-line parameters)
* Default UDP broadcast listening port: `9999` (modifiable via command-line parameters)

## Example Webpage Display

The frontend display page updates device information in real-time, including:

* Device Name
* IP Address
* MAC Address
* Device Model
* Software Version (linked to GitHub Release page)
* First Discovery Time
* Last Active Time

## Error Handling and Logging

The backend service provides detailed logs to quickly diagnose device discovery and communication issues.

## Dependencies

* github.com/gin-gonic/gin
* github.com/gorilla/websocket
* github.com/soulteary/gin-static

Use `go mod tidy` for automatic dependency management.

## License

This project uses the GPL-3.0 license.

---

Contributions are welcome! Feel free to submit code, raise issues, or open pull requests to help us continuously optimize and enhance CAN Finder!
