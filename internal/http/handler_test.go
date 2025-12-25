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

func setupTestServer() *TodoHandlers {
	store := todo.NewStore()
	service := todo.NewService(store)
	return NewTodoHandlers(service)
}

func TestHandler_CreateGetUpdateDeleteTodo(t *testing.T) {
	h := setupTestServer()

	//  CREATE
	todoData := todo.Todo{Title: "Test Task", Description: "Demo"}
	body, _ := json.Marshal(todoData)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(body))
	w := httptest.NewRecorder()
	h.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	var created todo.Todo
	if err := json.NewDecoder(w.Body).Decode(&created); err != nil {
		t.Fatal(err)
	}

	if created.ID == 0 {
		t.Fatal("expected non-zero ID for created todo")
	}

	//  GET BY ID
	req = httptest.NewRequest(http.MethodGet, "/todos/"+itoa(created.ID), nil)
	w = httptest.NewRecorder()
	h.GetByID(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var fetched todo.Todo
	if err := json.NewDecoder(w.Body).Decode(&fetched); err != nil {
		t.Fatal(err)
	}

	if fetched.Title != todoData.Title {
		t.Fatalf("expected title %s, got %s", todoData.Title, fetched.Title)
	}

	//  UPDATE
	updatedData := todo.Todo{Title: "Updated Task", Description: "Updated", Completed: true}
	body, _ = json.Marshal(updatedData)
	req = httptest.NewRequest(http.MethodPut, "/todos/"+itoa(created.ID), bytes.NewReader(body))
	w = httptest.NewRecorder()
	h.Update(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 on update, got %d", w.Code)
	}

	//  DELETE
	req = httptest.NewRequest(http.MethodDelete, "/todos/"+itoa(created.ID), nil)
	w = httptest.NewRecorder()
	h.Delete(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status 204 on delete, got %d", w.Code)
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
