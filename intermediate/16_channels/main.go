// 16_channels.go - Channels in Go

package main

import (
	"fmt"
	"time"
)

// ===== CHANNELS =====
// Channels are pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those
// values into another goroutine.

func worker(id int, jobs <-chan int, results chan<- int) {
	// jobs is receive-only channel (<-chan)
	// results is send-only channel (chan<-)

	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Millisecond * 100) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2 // Send result
	}
}

func main() {
	// ===== UNBUFFERED CHANNELS =====
	// Sending blocks until receiver is ready
	// Receiving blocks until sender is ready

	messages := make(chan string)

	go func() {
		fmt.Println("Goroutine sending message...")
		messages <- "ping" // Send
		fmt.Println("Goroutine sent message")
	}()

	fmt.Println("Main waiting for message...")
	msg := <-messages // Receive
	fmt.Println("Main received:", msg)

	// ===== BUFFERED CHANNELS =====
	// Buffered channels accept a limited number of values without
	// a corresponding receiver for those values.

	fmt.Println("\n--- Buffered Channel ---")
	buffered := make(chan string, 2) // Buffer size 2

	buffered <- "first"
	buffered <- "second"
	// buffered <- "third" // This would block if no receiver!

	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	// ===== CHANNEL SYNCHRONIZATION =====
	fmt.Println("\n--- Synchronization ---")
	done := make(chan bool)

	go func() {
		fmt.Print("Working...")
		time.Sleep(time.Second)
		fmt.Println("done")
		done <- true // Signal completion
	}()

	<-done // Block until notification received

	// ===== CHANNEL DIRECTIONS =====
	// Can specify if channel is for sending or receiving in function signature

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println("\nDirection result:", <-pongs)

	// ===== CLOSING CHANNELS & RANGE =====
	fmt.Println("\n--- Range over Channel ---")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // Close channel to indicate no more values

	// Range iterates until channel is closed
	for elem := range queue {
		fmt.Println(elem)
	}

	// ===== WORKER POOL PATTERN =====
	fmt.Println("\n--- Worker Pool ---")

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Signal no more jobs

	// Collect results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// Run: go run 16_channels.go
