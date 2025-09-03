package storage

import "todo-app/internal/models"

type TodoStorage interface {
    Create(todo *models.Todo) error
    GetByID(id string) (*models.Todo, error)
    GetAll() ([]*models.Todo, error)
    Update(todo *models.Todo) error
    Delete(id string) error
    FilterByStatus(status models.Status) ([]*models.Todo, error)
}