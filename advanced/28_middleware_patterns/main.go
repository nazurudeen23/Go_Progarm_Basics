// 28_middleware_patterns.go - Middleware Patterns

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ===== MIDDLEWARE PATTERN =====
// Middleware is a function that wraps an http.Handler or http.HandlerFunc
// to add functionality before/after the handler executes.

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ===== COMMON MIDDLEWARE PATTERNS =====

// 1. Logging Middleware
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		next(w, r)

		duration := time.Since(start)
		log.Printf("Request completed in %v", duration)
	}
}

// 2. Recovery Middleware (Panic recovery)
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next(w, r)
	}
}

// 3. CORS Middleware
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// 4. Rate Limiting Middleware (Simple)
func RateLimitMiddleware(requestsPerSecond int) Middleware {
	limiter := time.Tick(time.Second / time.Duration(requestsPerSecond))

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			<-limiter // Wait for rate limiter
			next(w, r)
		}
	}
}

// 5. Authentication Middleware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		// Validate token (simplified)
		if token != "Bearer valid-token" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

// ===== CHAINING MIDDLEWARE =====

// Chain multiple middleware
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		f = middlewares[i](f)
	}
	return f
}

// ===== HANDLERS =====

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API endpoint")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Something went wrong!")
}

func main() {
	// ===== USING MIDDLEWARE =====

	// Single middleware
	http.HandleFunc("/", LoggingMiddleware(homeHandler))

	// Multiple middleware (manual chaining)
	http.HandleFunc("/api",
		LoggingMiddleware(
			RecoveryMiddleware(
				CORSMiddleware(apiHandler),
			),
		),
	)

	// Using Chain helper
	http.HandleFunc("/protected",
		Chain(apiHandler,
			LoggingMiddleware,
			RecoveryMiddleware,
			AuthMiddleware,
		),
	)

	// Test panic recovery
	http.HandleFunc("/panic",
		Chain(panicHandler,
			LoggingMiddleware,
			RecoveryMiddleware,
		),
	)

	// Rate limited endpoint
	rateLimiter := RateLimitMiddleware(2) // 2 requests per second
	http.HandleFunc("/limited",
		Chain(apiHandler,
			LoggingMiddleware,
			rateLimiter,
		),
	)

	fmt.Println("Server starting on :8080")
	fmt.Println("\nEndpoints:")
	fmt.Println("  GET  /           - Home (with logging)")
	fmt.Println("  GET  /api        - API (with logging, recovery, CORS)")
	fmt.Println("  GET  /protected  - Protected (requires auth)")
	fmt.Println("  GET  /panic      - Panic test (with recovery)")
	fmt.Println("  GET  /limited    - Rate limited (2 req/sec)")

	fmt.Println("\nTest protected endpoint:")
	fmt.Println("  curl http://localhost:8080/protected -H 'Authorization: Bearer valid-token'")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
===== TESTING =====

# Home
curl http://localhost:8080/

# API
curl http://localhost:8080/api

# Protected (unauthorized)
curl http://localhost:8080/protected

# Protected (authorized)
curl http://localhost:8080/protected -H "Authorization: Bearer valid-token"

# Panic recovery
curl http://localhost:8080/panic

# Rate limiting (run multiple times quickly)
for i in {1..5}; do curl http://localhost:8080/limited & done
*/

// Run: go run 28_middleware_patterns.go
