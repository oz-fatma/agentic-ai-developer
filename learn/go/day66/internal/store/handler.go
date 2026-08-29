package store

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	Store *Store
}

func NewHandler(s *Store) http.Handler {
	h := &Handler{Store: s}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items", h.listItems)
	mux.HandleFunc("POST /items", h.createItem)
	mux.HandleFunc("GET /items/{id}", h.getItem)
	return mux
}

func (h *Handler) listItems(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.Store.List())
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	item := h.Store.Add(body.Title)
	writeJSON(w, item)
}

func (h *Handler) getItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/items/"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	item, ok := h.Store.Get(id)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	writeJSON(w, item)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
