package server

import (
	"afera-projects/internal/errors_pkg"
	"afera-projects/internal/model"
	"afera-projects/internal/repository"
	"afera-projects/internal/responses"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (s *Server) HandleGetProjectByID(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	log.Println("GET Project by id - /projects/{id}")

	id, err := uuid.Parse(mux.Vars(r)["id"])

	if err != nil {
		log.Printf("Invalid UUID : %s", err)

		responses.Response400(w, "Invalid UUID")

		return
	}

	project, err := s.ProjectRepository.GetByID(id)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println("Troubles while accessing database table (projects) width id")

		responses.Response404(w, "Project not found")

		return
	}

	if err != nil {
		log.Printf("Troubles while accessing database table (projects) width id: %s", err)

		responses.Response500(w, "Problems getting data from the database")

		return
	}

	responses.Response200(w, project)
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
