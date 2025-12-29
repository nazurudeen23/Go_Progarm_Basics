// 21_advanced_concurrency.go - Advanced Concurrency Patterns

package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== PATTERN 1: FAN-OUT, FAN-IN =====
// Fan-out: Multiple goroutines reading from the same channel
// Fan-in: Multiple goroutines writing to the same channel

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Close out when all output goroutines are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// ===== PATTERN 2: PIPELINE =====
func pipeline() {
	// Stage 1: Generate numbers
	nums := producer(1, 2, 3, 4, 5)

	// Stage 2: Square them
	squared := square(nums)

	// Stage 3: Consume
	for n := range squared {
		fmt.Println(n)
	}
}

// ===== PATTERN 3: WORKER POOL WITH RESULTS =====
type Job struct {
	ID   int
	Data int
}

type Result struct {
	Job    Job
	Result int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(100 * time.Millisecond)
		results <- Result{Job: job, Result: job.Data * 2}
	}
}

// ===== PATTERN 4: RATE LIMITING =====
func rateLimiter() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Rate limiter: 1 request per 200ms
	limiter := time.Tick(200 * time.Millisecond)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for req := range requests {
		<-limiter // Wait for limiter
		<-ticker.C // Wait for limiter
		fmt.Println("Request", req, time.Now())
	}
}

// ===== PATTERN 5: SEMAPHORE (LIMITING CONCURRENCY) =====
func semaphore() {
	const maxConcurrent = 3
	sem := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        // Acquire
			defer func() { <-sem }() // Release

			fmt.Printf("Task %d started\n", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Task %d finished\n", id)
		}(i)
	}

	wg.Wait()
}

func main() {
	// ===== PIPELINE DEMO =====
	fmt.Println("=== Pipeline ===")
	pipeline()

	// ===== FAN-OUT, FAN-IN DEMO =====
	fmt.Println("\n=== Fan-Out, Fan-In ===")
	in := producer(1, 2, 3, 4, 5, 6, 7, 8)

	// Fan-out: Multiple workers
	c1 := square(in)
	c2 := square(in)

	// Fan-in: Merge results
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

	// ===== WORKER POOL DEMO =====
	fmt.Println("\n=== Worker Pool ===")
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- Job{ID: j, Data: j * 10}
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Result: Job %d = %d\n", result.Job.ID, result.Result)
	}

	// ===== RATE LIMITING DEMO =====
	fmt.Println("\n=== Rate Limiting ===")
	rateLimiter()

	// ===== SEMAPHORE DEMO =====
	fmt.Println("\n=== Semaphore (Max 3 concurrent) ===")
	semaphore()
}

// Run: go run 21_advanced_concurrency.go
