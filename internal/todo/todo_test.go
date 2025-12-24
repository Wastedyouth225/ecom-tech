package todo

import (
	"errors"
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
	if !errors.Is(err, ErrInvalidTitle) {
		t.Fatalf("expected ErrInvalidTitle, got %v", err)
	}

	_, err = service.UpdateTodo(1, Todo{Title: ""})
	if !(errors.Is(err, ErrInvalidTitle) || errors.Is(err, ErrNotFound)) {
		t.Fatalf("expected ErrInvalidTitle or ErrNotFound, got %v", err)
	}
}

func TestUpdateTodo(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	created, _ := service.CreateTodo(Todo{Title: "Old", Description: "Desc"})

	updated, err := service.UpdateTodo(created.ID, Todo{Title: "New", Description: "Updated"})
	if err != nil {
		t.Fatal(err)
	}

	if updated.Title != "New" {
		t.Fatalf("expected title %s, got %s", "New", updated.Title)
	}
}

func TestDeleteTodo(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	created, _ := service.CreateTodo(Todo{Title: "ToDelete"})

	err := service.DeleteTodo(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.GetTodo(created.ID)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestNonExistentTodo(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	// Get
	_, err := service.GetTodo(999)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	// Update
	_, err = service.UpdateTodo(999, Todo{Title: "Nope"})
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	// Delete
	err = service.DeleteTodo(999)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
