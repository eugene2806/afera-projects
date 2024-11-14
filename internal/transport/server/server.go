package server

import (
	"log"
	"net/http"
)

type Server struct {
}

func BuildServer() *Server {
	return &Server{}
}

func (s *Server) GetPing(W http.ResponseWriter, r *http.Request) {
	log.Println("SUPER PING")
}
