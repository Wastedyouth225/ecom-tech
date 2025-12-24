package todo

import "errors"

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) CreateTodo(t Todo) (Todo, error) {
	if t.Title == "" {
		return Todo{}, ErrInvalidTitle
	}
	return s.store.Create(t), nil
}

func (s *Service) GetTodos() []Todo {
	return s.store.GetAll()
}

func (s *Service) GetTodo(id int) (Todo, error) {
	return s.store.GetByID(id)
}

func (s *Service) UpdateTodo(id int, t Todo) (Todo, error) {
	if t.Title == "" {
		return Todo{}, ErrInvalidTitle
	}
	return s.store.Update(id, t)
}

func (s *Service) DeleteTodo(id int) error {
	return s.store.Delete(id)
}

var ErrInvalidTitle = errors.New("title cannot be empty")
