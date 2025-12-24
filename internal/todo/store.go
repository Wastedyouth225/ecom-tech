package todo

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("todo not found")

type Store struct {
	mu     sync.RWMutex
	todos  map[int]Todo
	nextID int
}

func NewStore() *Store {
	return &Store{
		todos:  make(map[int]Todo),
		nextID: 1,
	}
}

func (s *Store) Create(t Todo) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.ID = s.nextID
	s.nextID++
	s.todos[t.ID] = t
	return t
}

func (s *Store) GetAll() []Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Todo, 0, len(s.todos))
	for _, t := range s.todos {
		result = append(result, t)
	}
	return result
}

func (s *Store) GetByID(id int) (Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, ok := s.todos[id]
	if !ok {
		return Todo{}, ErrNotFound
	}
	return t, nil
}

func (s *Store) Update(id int, t Todo) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return Todo{}, ErrNotFound
	}
	t.ID = id
	s.todos[id] = t
	return t, nil
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return ErrNotFound
	}
	delete(s.todos, id)
	return nil
}
