// 18_defer_panic_recover.go - Defer, Panic, and Recover

package main

import (
	"fmt"
	"os"
)

// ===== DEFER =====
// A defer statement defers the execution of a function until the surrounding
// function returns.
// Arguments are evaluated immediately, but the function call is not executed
// until the surrounding function returns.

func main() {
	// ===== DEFER BASICS =====
	fmt.Println("Start")
	defer fmt.Println("Deferred: End") // Executed last

	fmt.Println("Middle")

	// ===== STACKING DEFERS =====
	// Deferred calls are pushed onto a stack (LIFO)
	fmt.Println("\nCounting with defer:")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("\nDone counting (defers will print now in reverse)")

	// ===== PRACTICAL USE: CLEANUP =====
	// Defer is commonly used for cleanup (closing files, unlocking mutexes)
	fmt.Println("\n--- File Cleanup Example ---")
	createAndWriteFile("test.txt")

	// ===== PANIC AND RECOVER =====
	fmt.Println("\n--- Panic and Recover ---")

	// This function will panic but recover
	mayPanic()

	fmt.Println("Main continued after recovery")
}

func createAndWriteFile(filename string) {
	fmt.Println("Creating file...")
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Close will be called when function exits, even if error occurs later
	defer func() {
		fmt.Println("Closing file...")
		f.Close()
		os.Remove(filename) // Cleanup for this example
	}()

	fmt.Println("Writing to file...")
	fmt.Fprintln(f, "data")
}

func mayPanic() {
	// Recover must be called inside a deferred function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Calling panic...")
	panic("A problem occurred!")

	// fmt.Println("This will not execute")
}

// Run: go run 18_defer_panic_recover.go
