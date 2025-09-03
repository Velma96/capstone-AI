package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
	"todo-app/internal/models"
)

type JSONFileStorage struct {
	filepath string
	todos    map[string]*models.Todo
	mutex    sync.RWMutex
}

func NewJSONFileStorage(filepath string) (*JSONFileStorage, error) {
	storage := &JSONFileStorage{
		filepath: filepath,
		todos:    make(map[string]*models.Todo),
	}

	// Load existing data if file exists
	if _, err := os.Stat(filepath); err == nil {
		if err := storage.load(); err != nil {
			return nil, err
		}
	}

	return storage, nil
}

func (j *JSONFileStorage) load() error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	file, err := os.Open(j.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&j.todos)
}

func (j *JSONFileStorage) save() error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	file, err := os.Create(j.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(j.todos)
}

func (j *JSONFileStorage) Create(todo *models.Todo) error {
	j.mutex.Lock()
	j.todos[todo.ID] = todo
	j.mutex.Unlock()

	return j.save()
}

func (j *JSONFileStorage) GetByID(id string) (*models.Todo, error) {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	todo, exists := j.todos[id]
	if !exists {
		return nil, ErrNotFound
	}
	return todo, nil
}

func (j *JSONFileStorage) GetAll() ([]*models.Todo, error) {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	todos := make([]*models.Todo, 0, len(j.todos))
	for _, todo := range j.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (j *JSONFileStorage) Update(todo *models.Todo) error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if _, exists := j.todos[todo.ID]; !exists {
		return ErrNotFound
	}

	todo.UpdatedAt = time.Now()
	j.todos[todo.ID] = todo

	return j.save()
}

func (j *JSONFileStorage) Delete(id string) error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if _, exists := j.todos[id]; !exists {
		return ErrNotFound
	}

	delete(j.todos, id)
	return j.save()
}

func (j *JSONFileStorage) FilterByStatus(status models.Status) ([]*models.Todo, error) {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	var filtered []*models.Todo
	for _, todo := range j.todos {
		if todo.Status == status {
			filtered = append(filtered, todo)
		}
	}
	return filtered, nil
}