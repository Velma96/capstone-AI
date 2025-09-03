# Go (Golang) – A Beginner's Toolkit with Todo App Example

##  Quick Start

```bash
# Clone and setup
git clone <your-repo-url>
cd todo-app
go mod tidy
go run ./cmd/server
```

*In another terminal:*
```bash
cd todo-app
go build -o todo ./cmd/cli
./todo list
```

## 1. Title & Objective

**Title:** "Getting Started with Go (Golang) – A Beginner's Toolkit with Todo App"

- **Technology:** Go programming language (Golang)
- **Why Go?** Simple syntax, fast compilation, excellent concurrency support, perfect for web services
- **End Goal:** Build a complete Todo List application with REST API and CLI interface

## 2. What You'll Build

A full-featured Todo application with:
- ✅ REST API backend with Gorilla Mux
- ✅ Command-line interface (CLI)
- ✅ In-memory and JSON file storage options
- ✅ CRUD operations (Create, Read, Update, Delete)
- ✅ Filtering by status (pending, in progress, completed)

## 3. System Requirements

- **OS:** Linux, macOS, or Windows
- **Tools:** VS Code (recommended) or any text editor
- **Go Version:** 1.18+ 
- **Memory:** 2GB RAM minimum

## 4. Installation & Setup

### Install Go
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# macOS with Homebrew
brew install go

# Or download from: https://go.dev/dl/
```

### Verify Installation
```bash
go version
# Should show: go version go1.21.0 linux/amd64 (or similar)
```

### Create Project Workspace
```bash
mkdir -p go-projects/todo-app
cd go-projects/todo-app
```

## 5. AI Prompt Journal

### Prompt 1: 
**"I'm learning Go, coming from Python. Could you explain Go’s type system by comparing it to Python’s dynamic typing, especially around variables and function parameters?"**

### Prompt 2:
**"I’ve learned about goroutines in Go. Could you help me understand:

Why Go uses goroutines instead of OS threads?

What are the performance implications?

How would scaling to millions of goroutines compare to Python’s async/await or Java threads?"**


### Prompt 3:
**"Here’s how I understand Go’s slices:
‘A slice in Go is like a view of an array. It points to an underlying array, and changing the slice can change the original array. Slices also keep track of length and capacity.’

Could you check my explanation? What did I get right, and what am I missing?"**

## 6. Complete Setup Guide

### Step 1: Initialize Project
```bash
cd todo-app
go mod init todo-app
```

### Step 2: Install Dependencies
```bash
go get github.com/gorilla/mux
go get github.com/google/uuid
go mod tidy
```

### Step 3: Create Project Structure
```bash
mkdir -p cmd/server cmd/cli internal/models internal/storage internal/handlers internal/service pkg/utils
```

### Step 4: Add All Project Files
*Create all the files from the complete example provided in the previous response*

### Step 5: Test Build
```bash
cd todo-app
go build ./...
```

### Step 6: Run the Application
```bash
# Terminal 1 - Start server
cd todo-app
go run ./cmd/server

# Terminal 2 - Use CLI
cd todo-app
go build -o todo ./cmd/cli
./todo create
./todo list
```

## 7. Project Structure Explained

```
todo-app/
├── cmd/                 # Main applications
│   ├── server/         # HTTP server main
│   └── cli/            # Command-line interface
├── internal/           # Private application code
│   ├── models/         # Data structures (Todo)
│   ├── storage/        # Storage interfaces & implementations
│   ├── handlers/       # HTTP request handlers
│   └── service/        # Business logic layer
├── pkg/utils/          # Shared utilities
├── go.mod             # Module dependencies
└── go.sum             # Dependency checksums
```

## 8. Key Features

- **RESTful API** with proper HTTP methods
- **CLI Interface** for terminal usage
- **Multiple Storage** options (memory + JSON file)
- **Error Handling** with proper status codes
- **Filtering** by todo status
- **Structured Logging**

## 9. API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/todos` | List all todos |
| POST | `/api/v1/todos` | Create new todo |
| GET | `/api/v1/todos/{id}` | Get specific todo |
| PUT | `/api/v1/todos/{id}` | Update todo |
| DELETE | `/api/v1/todos/{id}` | Delete todo |
| GET | `/api/v1/todos/filter?status={status}` | Filter todos |

## 10. CLI Commands

```bash
./todo create          # Create new todo
./todo list           # List all todos
./todo get <id>       # Get specific todo
./todo update <id>    # Update todo
./todo delete <id>    # Delete todo
./todo filter <status> # Filter by status
```

## 11. Common Issues & Solutions

### Issue: `'go' not recognized`
**Solution:**
```bash
# Add Go to PATH (Linux/macOS)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Or reinstall Go
```

### Issue: `missing go.sum entry`
**Solution:**
```bash
cd todo-app
go mod tidy
```

### Issue: `undefined: time`
**Solution:** Add missing import to your Go file:
```go
import "time"
```

### Issue: `address already in use`
**Solution:** Use a different port:
```bash
go run ./cmd/server --port=8081
```

## 12. Learning Outcomes

By completing this project, you'll understand:

- ✅ Go project structure and organization
- ✅ HTTP server creation with Gorilla Mux
- ✅ CLI application development
- ✅ Struct usage for data modeling
- ✅ Error handling patterns in Go
- ✅ Package management with Go modules
- ✅ REST API design principles

## 13. Next Steps

1. **Add Database:** Integrate PostgreSQL or SQLite
2. **Authentication:** Add user login/signup
3. **Frontend:** Create web interface with HTML/Templates
4. **Testing:** Write unit and integration tests
5. **Dockerize:** Create Docker containers
6. **Deploy:** Host on cloud platform (AWS, GCP, Heroku)

## 14. Resources

- **Official Documentation:** https://golang.org/doc/
- **Go Tour:** https://tour.golang.org/
- **Go by Example:** https://gobyexample.com/
- **Gorilla Mux:** https://github.com/gorilla/mux

## 15. Support

If you get stuck:

1. Check all import statements
2. Run `go mod tidy`
3. Verify all files are in correct directories
4. Ensure you're in the `todo-app` directory when running commands

```bash
# Remember: Always start from the project root!
cd todo-app
go run ./cmd/server
```

## Licence

You now have a complete, working Todo application in Go! This project gives you an excellent foundation in Go programming and web development concepts.


