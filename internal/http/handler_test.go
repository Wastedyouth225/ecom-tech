package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"ecom-tech/internal/todo"
)

func setupHandler() *Handler {
	store := todo.NewStore()
	service := todo.NewService(store)
	return NewHandler(service)
}

func TestHandler_CreateTodo_Success(t *testing.T) {
	handler := setupHandler()
	todoData := map[string]interface{}{
		"title":       "Test Task",
		"description": "Test Description",
		"completed":   false,
	}
	body, _ := json.Marshal(todoData)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Router().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rr.Code)
	}

	var response todo.Todo
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.Title != "Test Task" {
		t.Errorf("Expected title 'Test Task', got '%s'", response.Title)
	}
}

func TestHandler_CreateTodo_EmptyTitle(t *testing.T) {
	handler := setupHandler()
	todoData := map[string]interface{}{
		"title":       "",
		"description": "Test",
	}
	body, _ := json.Marshal(todoData)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Router().ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rr.Code)
	}
}

func TestHandler_GetTodo_NotFound(t *testing.T) {
	handler := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/todos/999", nil)
	rr := httptest.NewRecorder()

	handler.Router().ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rr.Code)
	}
}

func TestHandler_UpdateTodo_Success(t *testing.T) {
	handler := setupHandler()
	created, _ := handler.service.CreateTodo(todo.Todo{Title: "Old", Description: "Desc"})

	updateData := map[string]interface{}{
		"title":       "Updated",
		"description": "New Desc",
		"completed":   true,
	}
	body, _ := json.Marshal(updateData)
	req := httptest.NewRequest(http.MethodPut, "/todos/"+strconv.Itoa(created.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Router().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rr.Code)
	}

	var updated todo.Todo
	if err := json.Unmarshal(rr.Body.Bytes(), &updated); err != nil {
		t.Fatal(err)
	}

	if updated.Title != "Updated" || !updated.Completed {
		t.Errorf("Update failed: %+v", updated)
	}
}

func TestHandler_DeleteTodo_Success(t *testing.T) {
	handler := setupHandler()
	created, _ := handler.service.CreateTodo(todo.Todo{Title: "ToDelete", Description: "Desc"})

	req := httptest.NewRequest(http.MethodDelete, "/todos/"+strconv.Itoa(created.ID), nil)
	rr := httptest.NewRecorder()

	handler.Router().ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("Expected status 204, got %d", rr.Code)
	}

	_, err := handler.service.GetTodo(created.ID)
	if err == nil {
		t.Errorf("Expected todo to be deleted")
	}
}
