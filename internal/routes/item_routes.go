package routes

import (
	"net/http"

	"github.com/DecBat/DecCollectionManager/internal/handlers"
)

func SetupItemRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("GET /items", handler.ListItems())
	mux.HandleFunc("POST /items", handler.CreateItem())
	mux.HandleFunc("GET /items/{id}", handler.GetItem())
}
