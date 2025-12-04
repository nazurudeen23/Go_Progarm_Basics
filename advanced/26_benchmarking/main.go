// 26_benchmarking.go - Benchmarking in Go

package main

import "fmt"

// ===== BENCHMARKING =====
/*
Benchmarks measure the performance of code.

Benchmark files must:
1. End with _test.go
2. Import "testing"
3. Have functions named BenchmarkXxx(b *testing.B)

Example benchmark file (26_benchmarking_test.go):

	package main

	import "testing"

	func Fibonacci(n int) int {
		if n <= 1 {
			return n
		}
		return Fibonacci(n-1) + Fibonacci(n-2)
	}

	func BenchmarkFibonacci10(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fibonacci(10)
		}
	}

	func BenchmarkFibonacci20(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fibonacci(20)
		}
	}

Run benchmarks:
	go test -bench=.
	go test -bench=Fibonacci10
	go test -bench=. -benchmem  # Include memory stats

===== BENCHMARK BEST PRACTICES =====

1. Reset timer if setup is needed:

	func BenchmarkWithSetup(b *testing.B) {
		// Setup
		data := generateLargeDataset()

		b.ResetTimer() // Don't count setup time

		for i := 0; i < b.N; i++ {
			processData(data)
		}
	}

2. Prevent compiler optimizations:

	var result int

	func BenchmarkCalculation(b *testing.B) {
		var r int
		for i := 0; i < b.N; i++ {
			r = expensiveCalculation()
		}
		result = r // Prevent optimization
	}

3. Run sub-benchmarks:

	func BenchmarkSizes(b *testing.B) {
		sizes := []int{10, 100, 1000, 10000}

		for _, size := range sizes {
			b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					process(size)
				}
			})
		}
	}

===== COMPARING BENCHMARKS =====

1. Save baseline:
	go test -bench=. > old.txt

2. Make changes

3. Run again:
	go test -bench=. > new.txt

4. Compare with benchstat:
	go install golang.org/x/perf/cmd/benchstat@latest
	benchstat old.txt new.txt

===== PROFILING =====

CPU profiling:
	go test -bench=. -cpuprofile=cpu.prof
	go tool pprof cpu.prof

Memory profiling:
	go test -bench=. -memprofile=mem.prof
	go tool pprof mem.prof
*/

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println("This file demonstrates benchmarking concepts.")
	fmt.Println("Create a file named '26_benchmarking_test.go' to write actual benchmarks.")
	fmt.Println("\nExample benchmark:")
	fmt.Println(`
package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(10)
	}
}
`)
	fmt.Println("\nRun with: go test -bench=.")

	// Demo the functions
	fmt.Println("\nFibonacci(10):", Fibonacci(10))
	fmt.Println("FibonacciIterative(10):", FibonacciIterative(10))
}

// To practice:
// 1. Create 26_benchmarking_test.go
// 2. Write benchmarks for Fibonacci and FibonacciIterative
// 3. Run: go test -bench=. -benchmem
// 4. Compare performance
