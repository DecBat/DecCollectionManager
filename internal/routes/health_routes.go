package routes

import (
	"net/http"

	"github.com/DecBat/DecCollectionManager/internal/handlers"
)

func SetupHealthRoute(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("/health", handler.HealthHandler())
}
