package main

import(
    "golang.org/x/net/websocket"
    "fmt"
)

type Server struct {
    connections map[*websocket.Conn]bool
}

func NewServer() *Server {
    return &Server{
        connections: make(map[*websocket.Conn]bool),
    }
}

func (s *Server) handleWS(ws *websocket.Conn) {
    fmt.println("new connection from client:", ws.RemoteAddr())
}

func main() {
}
