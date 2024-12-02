package server

import (
	"afera-projects/internal/errors_pkg"
	"afera-projects/internal/model"
	"afera-projects/internal/repository"
	"afera-projects/internal/responses"
	"encoding/json"
	"errors"
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
	log.Println("SUPER PING")
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("content-type", "application/json")
}

func (s *Server) HandleGetProjectsList(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) HandleGetProjectByID(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) HandleCreateProject(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	log.Println("POST Project - /projects")

	var req model.ProjectRequest

	json.NewDecoder(r.Body).Decode(&req)

	resp, err := s.ProjectRepository.Create(req)

	if errors.Is(err, errors_pkg.ErrInvalidRequest) {
		log.Printf("%s", err)

		responses.Response400(w, "Invalid request")

		return
	}

	if err != nil {
		log.Printf("Failed to create project :%s", err)

		responses.Response500(w, "Failed to create project")

		return
	}

	responses.Response201(w, resp)

}

func (s *Server) HandleUpdateProject(w http.ResponseWriter, r *http.Request) {}

func (s *Server) HandleDeleteProject(w http.ResponseWriter, r *http.Request) {}
