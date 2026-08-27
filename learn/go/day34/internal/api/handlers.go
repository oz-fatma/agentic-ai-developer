package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day34/internal/models"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day34/internal/store"
)

const maxTitleLen = 100

type Handler struct {
	store store.TaskStore
	mux   *http.ServeMux
}

func NewHandler(s store.TaskStore) *Handler {
	h := &Handler{store: s, mux: http.NewServeMux()}
	h.mux.HandleFunc("GET /tasks", h.listTasks)
	h.mux.HandleFunc("POST /tasks", h.createTask)
	h.mux.HandleFunc("GET /tasks/{id}", h.getTask)
	h.mux.HandleFunc("PUT /tasks/{id}", h.updateTask)
	h.mux.HandleFunc("DELETE /tasks/{id}", h.deleteTask)
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) listTasks(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.List())
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var in models.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "request body must be valid JSON")
		return
	}
	if err := validateTitle(in.Title); err != nil {
		writeError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}
	t := h.store.Create(strings.TrimSpace(in.Title))
	writeJSON(w, http.StatusCreated, t)
}

func (h *Handler) getTask(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_id", "task id must be a positive integer")
		return
	}
	t, ok := h.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "not_found", "task not found")
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_id", "task id must be a positive integer")
		return
	}
	var in models.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "request body must be valid JSON")
		return
	}
	if in.Title != nil {
		if err := validateTitle(*in.Title); err != nil {
			writeError(w, http.StatusBadRequest, "validation_error", err.Error())
			return
		}
		trimmed := strings.TrimSpace(*in.Title)
		in.Title = &trimmed
	}
	t, ok := h.store.Update(id, in.Title, in.Done)
	if !ok {
		writeError(w, http.StatusNotFound, "not_found", "task not found")
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_id", "task id must be a positive integer")
		return
	}
	if !h.store.Delete(id) {
		writeError(w, http.StatusNotFound, "not_found", "task not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func validateTitle(title string) error {
	trimmed := strings.TrimSpace(title)
	if trimmed == "" {
		return errEmptyTitle
	}
	if utf8.RuneCountInString(trimmed) > maxTitleLen {
		return errTitleTooLong
	}
	return nil
}

var (
	errEmptyTitle   = validationError("title is required")
	errTitleTooLong = validationError("title must be at most 100 characters")
)

type validationError string

func (e validationError) Error() string { return string(e) }

func parseID(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}
