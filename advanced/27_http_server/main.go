// 27_http_server.go - HTTP Server in Go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ===== BASIC HTTP SERVER =====

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

// Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// JSON response
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handle different HTTP methods
func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newUser.ID = len(users) + 1
		users = append(users, newUser)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ===== MIDDLEWARE =====

// Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next(w, r)

		log.Printf("Completed in %v", time.Since(start))
	}
}

// Auth middleware (simple example)
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func main() {
	// ===== SIMPLE SERVER =====
	// http.HandleFunc("/", helloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// ===== SERVER WITH MULTIPLE ROUTES =====

	// Basic routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/users", loggingMiddleware(usersHandler))
	http.HandleFunc("/api/users", loggingMiddleware(userHandler))

	// Protected route
	http.HandleFunc("/protected", loggingMiddleware(authMiddleware(protectedHandler)))

	// Static file server
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// ===== CUSTOM SERVER =====
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Server starting on :8080")
	fmt.Println("Routes:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /users")
	fmt.Println("  GET  /api/users")
	fmt.Println("  POST /api/users")
	fmt.Println("  GET  /protected (requires Authorization: secret-token)")

	log.Fatal(server.ListenAndServe())
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have access to protected content!")
}

/*
===== TESTING THE SERVER =====

1. Start the server:
   go run 27_http_server.go

2. Test with curl:

   # GET request
   curl http://localhost:8080/

   # Get users
   curl http://localhost:8080/users

   # POST new user
   curl -X POST http://localhost:8080/api/users \
     -H "Content-Type: application/json" \
     -d '{"name":"Charlie"}'

   # Protected route (unauthorized)
   curl http://localhost:8080/protected

   # Protected route (authorized)
   curl http://localhost:8080/protected \
     -H "Authorization: secret-token"

===== POPULAR FRAMEWORKS =====

While net/http is powerful, you might want to use frameworks for larger apps:

1. Gin: github.com/gin-gonic/gin
2. Echo: github.com/labstack/echo
3. Fiber: github.com/gofiber/fiber
4. Chi: github.com/go-chi/chi
*/

// Run: go run 27_http_server.go
