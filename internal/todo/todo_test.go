package todo

import (
	"testing"
)

func TestCreateAndGetTodo(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	todo := &Todo{Title: "Test", Description: "Demo"}
	created, err := service.Create(todo)
	if err != nil {
		t.Fatal(err)
	}

	if created.ID != 1 {
		t.Errorf("expected ID 1, got %d", created.ID)
	}

	got, err := service.GetByID(1)
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "Test" {
		t.Errorf("expected title 'Test', got '%s'", got.Title)
	}
}

func TestValidation(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	todo := &Todo{Title: "", Description: "No title"}
	_, err := service.Create(todo)
	if err == nil {
		t.Fatal("expected error for empty title")
	}
}

func TestUpdateAndDelete(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	todo := &Todo{Title: "Task", Description: "Desc"}
	created, _ := service.Create(todo)

	// Update
	updated, err := service.Update(created.ID, &Todo{Title: "Updated", Description: "New", Completed: true})
	if err != nil {
		t.Fatal(err)
	}
	if !updated.Completed {
		t.Fatal("expected completed true")
	}

	// Delete
	if err := service.Delete(created.ID); err != nil {
		t.Fatal(err)
	}

	_, err = service.GetByID(created.ID)
	if err != ErrNotFound {
		t.Fatal("expected ErrNotFound after delete")
	}
}
