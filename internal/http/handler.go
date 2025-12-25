package http

import (
	"ecom-tech/internal/todo"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type TodoHandlers struct {
	service *todo.Service
}

func NewTodoHandlers(s *todo.Service) *TodoHandlers {
	return &TodoHandlers{service: s}
}

func (h *TodoHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	todos := h.service.GetTodos()
	writeJSON(w, http.StatusOK, todos)
}

func (h *TodoHandlers) Create(w http.ResponseWriter, r *http.Request) {
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
	writeJSON(w, http.StatusCreated, created) // 201 Created
}

func (h *TodoHandlers) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	t, err := h.service.GetTodo(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (h *TodoHandlers) Update(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var t todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	updated, err := h.service.UpdateTodo(id, t)
	if err != nil {
		// Проверяем тип ошибки, чтобы вернуть правильный статус
		if err == todo.ErrInvalidTitle {
			writeError(w, http.StatusBadRequest, err.Error())
		} else {
			writeError(w, http.StatusNotFound, "todo not found")
		}
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *TodoHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteTodo(id); err != nil {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func extractID(r *http.Request) (int, error) {
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	// Благодаря паттерну "todos/{id}" в роутере, мы уверены в структуре
	idStr := segments[1]
	return strconv.Atoi(idStr)
}
