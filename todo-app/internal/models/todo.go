package models

import (
    "time"
    "github.com/google/uuid"
)

type Status string

const (
    StatusPending   Status = "pending"
    StatusInProgress Status = "in_progress"
    StatusCompleted Status = "completed"
)

type Todo struct {
    ID          string    `json:"id"`
    Title       string    `json:"title" validate:"required,min=1,max=255"`
    Description string    `json:"description,omitempty"`
    Status      Status    `json:"status"`
    DueDate     time.Time `json:"due_date,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func NewTodo(title, description string, dueDate time.Time) *Todo {
    now := time.Now()
    return &Todo{
        ID:          uuid.New().String(),
        Title:       title,
        Description: description,
        Status:      StatusPending,
        DueDate:     dueDate,
        CreatedAt:   now,
        UpdatedAt:   now,
    }
}