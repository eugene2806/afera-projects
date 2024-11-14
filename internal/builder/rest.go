package builder

import (
	"github.com/gorilla/mux"
	"net/http"
)

func BuildRestServer(Port string, Handler *mux.Router) *http.Server {
	return &http.Server{
		Addr:    ":" + Port,
		Handler: Handler,
	}
}
