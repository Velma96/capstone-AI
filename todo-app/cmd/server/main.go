package main

import (
	"flag"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"todo-app/internal/handlers"
	"todo-app/internal/service"
	"todo-app/internal/storage"
)

func main() {
	// Command line flags
	storageType := flag.String("storage", "memory", "Storage type: memory or json")
	jsonFile := flag.String("json-file", "todos.json", "JSON file path for json storage")
	port := flag.String("port", "8080", "Server port")
	flag.Parse()

	// Initialize storage
	var store storage.TodoStorage

	switch *storageType {
	case "json":
		var err error
		store, err = storage.NewJSONFileStorage(*jsonFile)
		if err != nil {
			log.Fatalf("Failed to create JSON storage: %v", err)
		}
		log.Printf("Using JSON file storage: %s", *jsonFile)
	default:
		store = storage.NewMemoryStorage()
		log.Println("Using in-memory storage")
	}

	// Initialize service and handlers
	service := service.NewTodoService(store)
	handler := handlers.NewTodoHandler(service)
	
	// Create router
	router := mux.NewRouter()
	
	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	api.HandleFunc("/todos", handler.GetAllTodos).Methods("GET")
	api.HandleFunc("/todos/filter", handler.FilterTodos).Methods("GET")
	api.HandleFunc("/todos/{id}", handler.GetTodo).Methods("GET")
	api.HandleFunc("/todos/{id}", handler.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", handler.DeleteTodo).Methods("DELETE")

	// Health check
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
	
	// Start server
	addr := ":" + *port
	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}