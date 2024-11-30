package builder

import (
	"afera-projects/internal/transport/server"
	"github.com/gorilla/mux"
	"net/http"
)

type HandlerBuilder struct {
	server *server.Server
}

func NewHandlerBuilder(server *server.Server) *HandlerBuilder {
	return &HandlerBuilder{
		server: server,
	}
}
func (h *HandlerBuilder) BuildHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", h.server.GetPing).Methods(http.MethodGet)
	router.HandleFunc("/projects", h.server.HandleGetProjectsList).Methods(http.MethodGet)
	router.HandleFunc("/projects/{id}", h.server.HandleGetProjectByID).Methods(http.MethodGet)
	router.HandleFunc("/projects", h.server.HandleCreateProject).Methods(http.MethodPost)
	router.HandleFunc("/projects/{id}", h.server.HandleUpdateProject).Methods(http.MethodPut)
	router.HandleFunc("/projects/{id}", h.server.HandleDeleteProject).Methods(http.MethodDelete)

	return router
}
