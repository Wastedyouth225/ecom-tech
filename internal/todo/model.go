package todo

import "errors"

var ErrInvalidTitle = errors.New("title cannot be empty")

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Validate проверка корректности
func (t *Todo) Validate() error {
	if t.Title == "" {
		return ErrInvalidTitle
	}
	return nil
}
