package api

import "net/http"

// Handler is the planned HTTP entry point for the tasks MVP.
// Day 31 scaffold: routes and store wiring land in later days.
type Handler struct {
	Mux *http.ServeMux
}

func NewHandler() *Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	return &Handler{Mux: mux}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Mux.ServeHTTP(w, r)
}
