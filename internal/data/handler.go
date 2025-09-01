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
		d.storage = append(d.storage, "new data")
		json.NewEncoder(w).Encode(map[string]string{"message": "added new data"})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (d *Handler) GetHandlerMap() map[string]http.HandlerFunc {
	handlerMap := make(map[string]http.HandlerFunc)
	handlerMap["/data"] = d.DataHandler
	return handlerMap
}
