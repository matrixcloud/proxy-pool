package server

import (
	"fmt"
	"net/http"

	"github.com/matrixcloud/proxy-pool/core"
	"github.com/rs/zerolog/log"
)

// Server provides rest api
type Server struct {
	pool *core.Pool
	port int
}

// NewServer creates a server
func NewServer(pool *core.Pool, port int) *Server {
	return &Server{
		pool: pool,
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
	proxies := s.pool.Get(1)

	if len(proxies) > 0 {
		w.Write([]byte(proxies[0].Addr))
	} else {
		w.Write([]byte("empty"))
	}
}

// Start starts api server
func (s *Server) Start() {

	http.HandleFunc("/", s.index)
	http.HandleFunc("/health", s.health)
	http.HandleFunc("/proxy", s.proxy)

	log.Info().Msgf("Server is listening on: http://localhost:%d", s.port)

	http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
