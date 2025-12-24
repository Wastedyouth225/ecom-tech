package todo

import "errors"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Validate testing
func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	return nil
}
