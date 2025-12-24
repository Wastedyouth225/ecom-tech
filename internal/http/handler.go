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

func (h *Handler) handleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		todos := h.service.GetTodos()
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		var t todo.Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		created, err := h.service.CreateTodo(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(created)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		t, err := h.service.GetTodo(id)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(t)
	case http.MethodPut:
		var t todo.Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated, err := h.service.UpdateTodo(id, t)
		if err != nil {
			if err == todo.ErrInvalidTitle {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, "not found", http.StatusNotFound)
			}
			return
		}
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		err := h.service.DeleteTodo(id)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
