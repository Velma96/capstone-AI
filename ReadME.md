

# Getting Started with Go (Golang) – A Beginner’s Toolkit



## 1. Title & Objective

**Title:** "Getting Started with Go (Golang) – A Beginner’s Toolkit"

- **What technology did you choose?** I chose the Go programming language (often referred to as Golang).
- **Why did you choose it?** Go is a statically typed, compiled language designed by Google for simplicity, reliability, and efficiency. It's known for its concise syntax, fast compilation times, and excellent built-in support for concurrency and web services. It's a fantastic language for a beginner to learn as it avoids much of the complexity of older languages while being incredibly powerful.
- **What’s the end goal?** The end goal is to write, compile, and run a simple "Hello, World!" application, and then extend it into a simple HTTP server that responds to web requests.



## 2. Quick Summary of the Technology

- **What is it?** Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.
- **Where is it used?** It is used heavily for cloud-native development, web servers, DevOps tooling, command-line tools, and distributed systems. Famous projects built with Go include Docker, Kubernetes, Terraform, and Caddy.
- **One real-world example.** When you use a service like Netflix or YouTube, many of the backend services that manage traffic and serve video content are written in Go because of its performance and efficiency in handling millions of concurrent connections.


## 3. System Requirements

- **OS:** Linux, macOS, or Windows
- **Tools/Editors required:** A text editor (like VS Code) and the Go compiler.
- **Any packages:** The Go toolchain includes everything needed. No external package managers are required for this basic example.

---

## 4. Installation & Setup Instructions

### Step 1: Download and Install Go
1. Visit the official Go downloads page: [https://go.dev/dl/](https://go.dev/dl/)
2. Download the installer for your operating system (`.msi` for Windows, `.pkg` for macOS, `.tar.gz` for Linux).
3. Run the installer and follow the instructions. The default settings are usually fine.

### Step 2: Verify the Installation
```bash
go version
````

Expected output (example):

```
go version go1.25.0 linux/amd64
```

### Step 3: Set Up Your Workspace (Optional)

Organize your projects in one place:

```bash
mkdir go-projects
cd go-projects
```

---

## 5. Minimal Working Example

### Example 1: Hello, World!

```go
// hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World from Go! 🚀")
}
```

Run it:

```bash
go run hello.go
```

Output:

```
Hello, World from Go! 🚀
```

---

### Example 2: Simple HTTP Server

```go
// server.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my first Go server! 🌐")
    })

    fmt.Println("Server starting on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}
```

Run it:

```bash
go run server.go
```

Open in browser: [http://localhost:8080](http://localhost:8080)
Stop with `Ctrl + C`.

---

## 6. AI Prompt Journal

**Prompt:** "I am a complete beginner. Give me a step-by-step guide to writing and running a 'Hello, World' program in Go."

* **AI’s response:** Provided a clear, step-by-step guide with code and commands.
* **Evaluation:** Extremely helpful. It acted as a patient tutor and reduced beginner friction.

**Prompt:** "How do I create a basic HTTP server in Go that listens on port 8080 and returns a text response?"

* **AI’s response:** Provided the full server code and explained the imports, handler, and `ListenAndServe`.
* **Evaluation:** Very good. Explained key concepts like `http.ResponseWriter` clearly.

---

## 7. Common Issues & Fixes

* **Error:** `'go' not recognized`

  * *Cause:* PATH not set.
  * *Fix:* Add Go bin directory to PATH.

* **Error:** `imported and not used`

  * *Cause:* Unused import.
  * *Fix:* Remove or use the import.

* **Error:** `panic: listen tcp :8080: address already in use`

  * *Cause:* Port conflict.
  * *Fix:* Use another port (`:8081`).

* **Error:** `go: cannot find main module`

  * *Cause:* Wrong directory.
  * *Fix:* Run inside the project folder.

---

## 8. References

* [Go Official Website](https://go.dev/)
* [A Tour of Go](https://go.dev/tour/)
* [Go Standard Library](https://pkg.go.dev/std)
* [Go by Example](https://gobyexample.com/)
* [Hitchhiker's Guide to Go](https://github.com/beyondns/gotips)

Community:

* [Go Forum](https://forum.golangbridge.org/)
* [r/golang on Reddit](https://www.reddit.com/r/golang/)

---

## Working Codebase

GitHub repo (placeholder):[https://github.com/Velma96/capstone-AI]

Contents:

* `hello.go`
* `server.go`
* `README.md`

### Run Instructions

```bash
# Run Hello World
go run hello.go

# Run HTTP Server
go run server.go
```

Then open [http://localhost:8080](http://localhost:8080).

```

---

Do you want me to also add the **Go installation troubleshooting (Linux PATH setup guide)** into the README, so beginners won’t run into the same “`go not found`” problem you faced?
```
