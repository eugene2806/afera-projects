package server

import (
	"afera-projects/internal/errors_pkg"
	"afera-projects/internal/repository"
	"afera-projects/internal/responses"
	"errors"
	"log"
	"net/http"
	"strconv"
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
	initHeaders(w)

	log.Println("GET Project List - /projects?page=x&limit=y")

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	var numErr *strconv.NumError

	projects, fullCount, fullPage, err := s.ProjectRepository.GetAllProjects(page, limit)

	if errors.As(err, &numErr) {
		log.Printf("problem when passing parameters : %s", numErr.Err)

		responses.Response400(w, "invalid request parameters")

		return
	}

	if errors.Is(err, errors_pkg.ErrLessZero) {
		log.Printf("problem when passing parameters : %s", err)

		responses.Response400(w, "invalid request parameters")

		return
	}

	if err != nil {
		log.Printf("problem when passing parameters : %s", err)

		responses.Response500(w, "Problems getting data from the database")

		return
	}

	responses.ResponseProjects200(w, projects, fullCount, fullPage)
}
