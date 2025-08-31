// server.go
package main

// Import multiple packages required for HTTP functionality
import (
    "fmt"   // For formatting text
    "net/http" // For HTTP server capabilities
)

// The main function
func main() {
    // Define the route "/" and tell the server what function to call when this route is hit
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Fprintf writes to the HTTP response (w)
        fmt.Fprintf(w, "Welcome to my first Go server! 🌐")
    })

    // Print a message to the console so we know the server is starting
    fmt.Println("Server starting on http://localhost:8080")
    
    // Start the server on port 8080 and handle any errors
    // If ListenAndServe returns an error, we log it and exit
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err) // panic is used here for simplicity in this example
    }
}