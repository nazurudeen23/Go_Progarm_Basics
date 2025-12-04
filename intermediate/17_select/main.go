// 17_select.go - Select Statement

package main

import (
	"fmt"
	"time"
)

// ===== SELECT =====
// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case.
// If multiple are ready, it chooses one at random.

func main() {
	// ===== BASIC SELECT =====
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	fmt.Println("Waiting for channels...")

	// We'll receive from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from c2:", msg2)
		}
	}

	// ===== TIMEOUTS =====
	// Select is useful for implementing timeouts

	fmt.Println("\n--- Timeouts ---")
	c3 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "result"
	}()

	select {
	case res := <-c3:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Operation took too long.")
	}

	// ===== NON-BLOCKING OPERATIONS =====
	// Use default case for non-blocking sends/receives

	fmt.Println("\n--- Non-Blocking Operations ---")
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	// Non-blocking send
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent (no receiver ready)")
	}

	// Multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity")
	}
}

// Run: go run 17_select.go
