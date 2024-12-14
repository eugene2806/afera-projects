package server

import (
	"afera-projects/internal/repository"
	"log"
	"net/http"
)

type Server struct {
	ProjectRepository *repository.ProjectRepository
}

func BuildServer(rep *repository.ProjectRepository) *Server {
	return &Server{
		ProjectRepository: rep,
	}
}

func (s *Server) GetPing(W http.ResponseWriter, r *http.Request) {
	log.Println("SUPER PINGs")
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("content-type", "application/json")
}
