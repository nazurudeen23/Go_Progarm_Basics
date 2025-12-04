// 15_goroutines.go - Concurrency with Goroutines

package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== GOROUTINES =====
// A goroutine is a lightweight thread managed by the Go runtime.
// Started with the 'go' keyword.

func printNumbers(prefix string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s: %d\n", prefix, i)
	}
}

func main() {
	// ===== BASIC GOROUTINE =====
	fmt.Println("Starting goroutines...")

	// Start a new goroutine
	go printNumbers("Goroutine")

	// Run in main goroutine (synchronously)
	printNumbers("Main")

	// Wait a bit to ensure goroutine finishes (naive approach)
	time.Sleep(1 * time.Second)

	// ===== WAITGROUPS =====
	// The proper way to wait for goroutines is sync.WaitGroup

	fmt.Println("\n--- Using WaitGroup ---")
	var wg sync.WaitGroup

	// Add 2 to the counter (we will start 2 goroutines)
	wg.Add(2)

	// Anonymous function as goroutine
	go func() {
		defer wg.Done() // Decrement counter when done
		printNumbers("Worker 1")
	}()

	go func() {
		defer wg.Done()
		printNumbers("Worker 2")
	}()

	fmt.Println("Waiting for workers...")
	wg.Wait() // Block until counter is 0
	fmt.Println("All workers done!")

	// ===== RACE CONDITIONS =====
	// When multiple goroutines access shared data without synchronization

	fmt.Println("\n--- Race Condition Example (Fixed with Mutex) ---")

	counter := 0
	var mu sync.Mutex // Mutex protects the counter
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()

			// Critical section - lock before accessing shared resource
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg2.Wait()
	fmt.Printf("Final Counter: %d (Expected 1000)\n", counter)
}

// Run: go run 15_goroutines.go
