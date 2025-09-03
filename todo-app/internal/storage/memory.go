package storage

import (
	"errors"
	"sync"
	"time"
	"todo-app/internal/models"
)

type MemoryStorage struct {
	todos map[string]*models.Todo
	mutex sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[string]*models.Todo),
	}
}

func (m *MemoryStorage) Create(todo *models.Todo) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	m.todos[todo.ID] = todo
	return nil
}

func (m *MemoryStorage) GetByID(id string) (*models.Todo, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	
	todo, exists := m.todos[id]
	if !exists {
		return nil, ErrNotFound
	}
	return todo, nil
}

func (m *MemoryStorage) GetAll() ([]*models.Todo, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	
	todos := make([]*models.Todo, 0, len(m.todos))
	for _, todo := range m.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (m *MemoryStorage) Update(todo *models.Todo) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	if _, exists := m.todos[todo.ID]; !exists {
		return ErrNotFound
	}
	
	todo.UpdatedAt = time.Now()
	m.todos[todo.ID] = todo
	return nil
}

func (m *MemoryStorage) Delete(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	if _, exists := m.todos[id]; !exists {
		return ErrNotFound
	}
	
	delete(m.todos, id)
	return nil
}

func (m *MemoryStorage) FilterByStatus(status models.Status) ([]*models.Todo, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	
	var filtered []*models.Todo
	for _, todo := range m.todos {
		if todo.Status == status {
			filtered = append(filtered, todo)
		}
	}
	return filtered, nil
}

var ErrNotFound = errors.New("todo not found")