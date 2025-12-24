package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"ecom-tech/internal/todo"
)

type Handler struct {
	service *todo.Service
}

func NewHandler(s *todo.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", h.handleTodos)
	mux.HandleFunc("/todos/", h.handleTodoByID)
	return mux
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func (h *Handler) handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		todos := h.service.GetTodos()
		writeJSON(w, http.StatusOK, todos)
	case http.MethodPost:
		var t todo.Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		created, err := h.service.CreateTodo(t)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, created)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *Handler) handleTodoByID(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(segments) != 2 || segments[0] != "todos" {
		writeError(w, http.StatusBadRequest, "invalid path")
		return
	}
	id, err := strconv.Atoi(segments[1])
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		t, err := h.service.GetTodo(id)
		if err != nil {
			writeError(w, http.StatusNotFound, "not found")
			return
		}
		writeJSON(w, http.StatusOK, t)
	case http.MethodPut:
		var t todo.Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		updated, err := h.service.UpdateTodo(id, t)
		if err != nil {
			if err == todo.ErrInvalidTitle {
				writeError(w, http.StatusBadRequest, err.Error())
			} else {
				writeError(w, http.StatusNotFound, "not found")
			}
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		err := h.service.DeleteTodo(id)
		if err != nil {
			writeError(w, http.StatusNotFound, "not found")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
