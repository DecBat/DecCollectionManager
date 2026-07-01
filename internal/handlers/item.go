package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DecBat/DecCollectionManager/internal/store"
	"github.com/DecBat/DecCollectionManager/models"
)

func (h *Handler) ListItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := h.queries.ListItems(r.Context())
		if err != nil {
			http.Error(w, `{"error":"failed to list items"}`, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ToItemResponses(items))
	}
}

func (h *Handler) GetItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
			return
		}
		item, err := h.queries.GetItem(r.Context(), int32(id))
		if err != nil {
			http.Error(w, `{"error":"item not found"}`, http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ToItemResponse(item))
	}
}

func (h *Handler) CreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.CreateItemRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
			return
		}

		params := store.CreateItemParams{
			Title: req.Title,
			Price: req.Price,
		}
		if req.UserID != nil {
			params.UserID = sql.NullInt32{Int32: *req.UserID, Valid: true}
		}

		item, err := h.queries.CreateItem(r.Context(), params)
		if err != nil {
			http.Error(w, `{"error":"failed to create item"}`, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(models.ToItemResponse(item))
	}
}
