package todo

import "errors"

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

// Создает задачу с валидацией
func (s *Service) CreateTodo(t Todo) (Todo, error) {
	if err := t.Validate(); err != nil {
		return Todo{}, err
	}
	return s.store.Create(t), nil
}

// Получает все задачи
func (s *Service) GetTodos() []Todo {
	return s.store.GetAll()
}

// Получает задачу по ID
func (s *Service) GetTodo(id int) (Todo, error) {
	return s.store.GetByID(id)
}

// Обновляет задачу
func (s *Service) UpdateTodo(id int, t Todo) (Todo, error) {
	if err := t.Validate(); err != nil {
		return Todo{}, err
	}
	return s.store.Update(id, t)
}

// Удаляет задачу
func (s *Service) DeleteTodo(id int) error {
	return s.store.Delete(id)
}

// ErrInvalidTitle можно оставить для совместимости, но теперь Validate используется
var ErrInvalidTitle = errors.New("title cannot be empty")
