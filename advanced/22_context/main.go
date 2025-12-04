// 22_context.go - Context Package

package main

import (
	"context"
	"fmt"
	"time"
)

// ===== CONTEXT =====
// Context carries deadlines, cancellation signals, and request-scoped values
// across API boundaries and between processes.

// Use cases:
// 1. Cancellation propagation
// 2. Timeouts
// 3. Request-scoped values (e.g., request ID, user info)

func doWork(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: Cancelled! Reason: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s: Working...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// ===== 1. CONTEXT WITH CANCEL =====
	fmt.Println("=== Context with Cancel ===")

	ctx1, cancel1 := context.WithCancel(context.Background())

	go doWork(ctx1, "Worker 1")

	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling Worker 1...")
	cancel1() // Signal cancellation

	time.Sleep(500 * time.Millisecond)

	// ===== 2. CONTEXT WITH TIMEOUT =====
	fmt.Println("\n=== Context with Timeout ===")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2() // Always call cancel to release resources

	go doWork(ctx2, "Worker 2")

	time.Sleep(3 * time.Second) // Wait longer than timeout

	// ===== 3. CONTEXT WITH DEADLINE =====
	fmt.Println("\n=== Context with Deadline ===")

	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()

	go doWork(ctx3, "Worker 3")

	time.Sleep(2 * time.Second)

	// ===== 4. CONTEXT WITH VALUE =====
	fmt.Println("\n=== Context with Value ===")

	type key string
	const userKey key = "user"

	ctx4 := context.WithValue(context.Background(), userKey, "Alice")

	processRequest(ctx4)

	// ===== 5. PRACTICAL EXAMPLE: HTTP REQUEST WITH TIMEOUT =====
	fmt.Println("\n=== Simulated HTTP Request ===")

	ctx5, cancel5 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel5()

	result := make(chan string, 1)

	go func() {
		// Simulate slow API call
		time.Sleep(2 * time.Second)
		result <- "API Response"
	}()

	select {
	case <-ctx5.Done():
		fmt.Println("Request timeout:", ctx5.Err())
	case res := <-result:
		fmt.Println("Got result:", res)
	}
}

func processRequest(ctx context.Context) {
	type key string
	const userKey key = "user"

	if user, ok := ctx.Value(userKey).(string); ok {
		fmt.Printf("Processing request for user: %s\n", user)
	} else {
		fmt.Println("No user in context")
	}
}

// Run: go run 22_context.go
