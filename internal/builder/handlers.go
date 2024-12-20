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

	return router
}
