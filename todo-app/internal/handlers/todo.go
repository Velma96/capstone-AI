package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    
    "github.com/gorilla/mux"
    "todo-app/internal/models"
    "todo-app/internal/service"
	"todo-app/internal/storage"
)

type TodoHandler struct {
    service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
    return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Title       string    `json:"title"`
        Description string    `json:"description"`
        DueDate     time.Time `json:"due_date"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    todo, err := h.service.CreateTodo(request.Title, request.Description, request.DueDate)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    todo, err := h.service.GetTodo(id)
    if err != nil {
        if err == storage.ErrNotFound {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := h.service.GetAllTodos()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    var request struct {
        Title       string        `json:"title"`
        Description string        `json:"description"`
        Status      models.Status `json:"status"`
        DueDate     time.Time     `json:"due_date"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    todo, err := h.service.UpdateTodo(id, request.Title, request.Description, request.Status, request.DueDate)
    if err != nil {
        if err == storage.ErrNotFound {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    if err := h.service.DeleteTodo(id); err != nil {
        if err == storage.ErrNotFound {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}

func (h *TodoHandler) FilterTodos(w http.ResponseWriter, r *http.Request) {
    status := r.URL.Query().Get("status")
    if status == "" {
        http.Error(w, "Status parameter is required", http.StatusBadRequest)
        return
    }
    
    todos, err := h.service.FilterByStatus(models.Status(status))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}