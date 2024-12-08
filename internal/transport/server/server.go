package server

import (
	"afera-projects/internal/repository"
	"afera-projects/internal/responses"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
