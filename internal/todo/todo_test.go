package todo

import (
	"testing"
)

func TestCreateAndGetTodo(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	todo := Todo{Title: "Test", Description: "Demo"}
	created, err := service.CreateTodo(todo)
	if err != nil {
		t.Fatal(err)
	}

	fetched, err := service.GetTodo(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	if fetched.Title != todo.Title {
		t.Fatalf("expected title %s, got %s", todo.Title, fetched.Title)
	}
}

func TestValidation(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	_, err := service.CreateTodo(Todo{Title: ""})
	if err != ErrInvalidTitle {
		t.Fatalf("expected ErrInvalidTitle, got %v", err)
	}
}
