package handlers

import "github.com/DecBat/DecCollectionManager/internal/store"

type Handler struct {
	// DB instance
	// Query stores
	queries *store.Queries
}

func NewHandlers(queries *store.Queries) *Handler {
	return &Handler{queries: queries}
}
