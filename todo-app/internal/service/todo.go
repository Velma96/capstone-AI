package service

import (
	"time"
	"todo-app/internal/models"
	"todo-app/internal/storage"
)

type TodoService struct {
	storage storage.TodoStorage
}

func NewTodoService(storage storage.TodoStorage) *TodoService {
	return &TodoService{storage: storage}
}

func (s *TodoService) CreateTodo(title, description string, dueDate time.Time) (*models.Todo, error) {
	todo := models.NewTodo(title, description, dueDate)
	if err := s.storage.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) GetTodo(id string) (*models.Todo, error) {
	return s.storage.GetByID(id)
}

func (s *TodoService) GetAllTodos() ([]*models.Todo, error) {
	return s.storage.GetAll()
}

func (s *TodoService) UpdateTodo(id, title, description string, status models.Status, dueDate time.Time) (*models.Todo, error) {
	todo, err := s.storage.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	if title != "" {
		todo.Title = title
	}
	if description != "" {
		todo.Description = description
	}
	if status != "" {
		todo.Status = status
	}
	if !dueDate.IsZero() {
		todo.DueDate = dueDate
	}
	
	if err := s.storage.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) DeleteTodo(id string) error {
	return s.storage.Delete(id)
}

func (s *TodoService) FilterByStatus(status models.Status) ([]*models.Todo, error) {
	return s.storage.FilterByStatus(status)
}