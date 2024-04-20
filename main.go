package main

import(
    "fmt"
    "io"
    "net/http"
    "golang.org/x/net/websocket"
)

type Server struct {
    conn map[*websocket.Conn]bool
}

func NewServer() *Server {
    return &Server{
        conn: make(map[*websocket.Conn]bool),
    }
}

func (s *Server) handleWS(ws *websocket.Conn) {
    fmt.Println("new connection from client:", ws.RemoteAddr())

    s.conn[ws] = true

    s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
    buf := make([]byte, 1024)
    for {
        n, err := ws.Read(buf)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("read error:", err)
            continue
        }
        msg := buf[:n]
        fmt.Println(string(msg))
        ws.Write([]byte("message sent"))
    }
}

func main() {
    server := NewServer()
    http.Handle("/ws", websocket.Handler(server.handleWS))
    http.ListenAndServe(":4444", nil)
}

