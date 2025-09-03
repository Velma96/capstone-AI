package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
)

type Todo struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    DueDate     time.Time `json:"due_date"`
    CreatedAt   time.Time `json:"created_at"`
}

const baseURL = "http://localhost:8080/api/v1"

func main() {
    if len(os.Args) < 2 {
        printUsage()
        return
    }
    
    switch os.Args[1] {
    case "create":
        createTodo()
    case "list":
        listTodos()
    case "get":
        if len(os.Args) < 3 {
            fmt.Println("Please provide todo ID")
            return
        }
        getTodo(os.Args[2])
    case "update":
        if len(os.Args) < 3 {
            fmt.Println("Please provide todo ID")
            return
        }
        updateTodo(os.Args[2])
    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Please provide todo ID")
            return
        }
        deleteTodo(os.Args[2])
    case "filter":
        if len(os.Args) < 3 {
            fmt.Println("Please provide status (pending/in_progress/completed)")
            return
        }
        filterTodos(os.Args[2])
    default:
        printUsage()
    }
}

func printUsage() {
    fmt.Println(`Todo CLI Usage:
  create - Create a new todo
  list - List all todos
  get <id> - Get a specific todo
  update <id> - Update a todo
  delete <id> - Delete a todo
  filter <status> - Filter todos by status`)
}

func createTodo() {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Title: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)
    
    fmt.Print("Description: ")
    description, _ := reader.ReadString('\n')
    description = strings.TrimSpace(description)
    
    fmt.Print("Due Date (YYYY-MM-DD): ")
    dateStr, _ := reader.ReadString('\n')
    dateStr = strings.TrimSpace(dateStr)
    
    var dueDate time.Time
    if dateStr != "" {
        var err error
        dueDate, err = time.Parse("2006-01-02", dateStr)
        if err != nil {
            fmt.Println("Invalid date format")
            return
        }
    }
    
    data := map[string]interface{}{
        "title":       title,
        "description": description,
        "due_date":    dueDate,
    }
    
    jsonData, _ := json.Marshal(data)
    resp, err := http.Post(baseURL+"/todos", "application/json", strings.NewReader(string(jsonData)))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    
    var todo Todo
    json.NewDecoder(resp.Body).Decode(&todo)
    printTodo(&todo)
}

func listTodos() {
    resp, err := http.Get(baseURL + "/todos")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    
    var todos []Todo
    json.NewDecoder(resp.Body).Decode(&todos)
    
    for _, todo := range todos {
        printTodo(&todo)
        fmt.Println("---")
    }
}

func getTodo(id string) {
    resp, err := http.Get(baseURL + "/todos/" + id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    
    if resp.StatusCode == 404 {
        fmt.Println("Todo not found")
        return
    }
    
    var todo Todo
    json.NewDecoder(resp.Body).Decode(&todo)
    printTodo(&todo)
}

func updateTodo(id string) {
    // Implementation similar to create but with PUT
    fmt.Println("Update functionality would go here")
}

func deleteTodo(id string) {
    client := &http.Client{}
    req, _ := http.NewRequest("DELETE", baseURL+"/todos/"+id, nil)
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    
    if resp.StatusCode == 204 {
        fmt.Println("Todo deleted successfully")
    } else {
        fmt.Println("Error deleting todo")
    }
}

func filterTodos(status string) {
    resp, err := http.Get(baseURL + "/todos/filter?status=" + status)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    
    var todos []Todo
    json.NewDecoder(resp.Body).Decode(&todos)
    
    for _, todo := range todos {
        printTodo(&todo)
        fmt.Println("---")
    }
}

func printTodo(todo *Todo) {
    fmt.Printf("ID: %s\n", todo.ID)
    fmt.Printf("Title: %s\n", todo.Title)
    fmt.Printf("Description: %s\n", todo.Description)
    fmt.Printf("Status: %s\n", todo.Status)
    fmt.Printf("Due Date: %s\n", todo.DueDate.Format("2006-01-02"))
    fmt.Printf("Created At: %s\n", todo.CreatedAt.Format("2006-01-02 15:04:05"))
}