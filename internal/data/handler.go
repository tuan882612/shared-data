package data

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	storage []string
}

func NewDataHandler() *Handler {
	return &Handler{
		storage: make([]string, 0),
	}
}

func (d *Handler) DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(map[string]any{
			"storage": d.storage,
		})
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "added new data"})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
