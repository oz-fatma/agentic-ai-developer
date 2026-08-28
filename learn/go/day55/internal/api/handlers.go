package api

import (
	"encoding/json"
	"net/http"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day55/internal/auth"
)

type Handler struct {
	auth *auth.Service
}

func NewHandler(svc *auth.Service) http.Handler {
	h := &Handler{auth: svc}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", h.register)
	mux.HandleFunc("POST /login", h.login)
	mux.HandleFunc("GET /me", h.me)
	return mux
}

type credsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var req credsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	user, err := h.auth.Register(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"id": user.ID, "email": user.Email})
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var req credsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	token, err := h.auth.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) me(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	claims, err := h.auth.ParseToken(authHeader[7:])
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"email": claims.Subject, "role": claims.Role})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
