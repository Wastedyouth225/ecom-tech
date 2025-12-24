package todo

import "errors"

// Todo представляет задачу
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Validate проверяет корректность полей задачи
func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	return nil
}
