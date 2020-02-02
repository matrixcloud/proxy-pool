package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matrixcloud/proxy-pool/db"
)

// Server provides rest api
type Server struct {
	conn *db.Client
	port int
}

// NewServer creates a server
func NewServer(conn *db.Client, port int) *Server {
	return &Server{
		conn: conn,
		port: port,
	}
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is started"))
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "up"}`))
}

func (s *Server) proxy(w http.ResponseWriter, r *http.Request) {
	proxies := s.conn.Get(1)

	w.Write([]byte(proxies[0]))
}

// Start starts api server
func (s *Server) Start() {
	http.HandleFunc("/", s.index)
	http.HandleFunc("/health", s.health)
	http.HandleFunc("/proxy", s.proxy)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil))
}
