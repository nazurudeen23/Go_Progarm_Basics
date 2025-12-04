// 13_error_handling.go - Error Handling in Go

package main

import (
	"errors"
	"fmt"
	"os"
)

// ===== BASIC ERROR HANDLING =====
// Go doesn't have try/catch exceptions.
// Functions return errors as values (usually the last return value).

// Function returning an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Create a new error with a message
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// ===== CUSTOM ERRORS =====
// Any type implementing the Error() string method is an error

type NegativeSqrtError struct {
	Value float64
}

func (e *NegativeSqrtError) Error() string {
	return fmt.Sprintf("cannot calculate square root of negative number: %v", e.Value)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// Return pointer to custom error struct
		return 0, &NegativeSqrtError{Value: x}
	}
	// Simplified implementation
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// ===== WRAPPING ERRORS =====
// Go 1.13+ introduced error wrapping
func openFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		// %w verb wraps the error
		return fmt.Errorf("failed to open config: %w", err)
	}
	return nil
}

func main() {
	// ===== CHECKING ERRORS =====
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// Trigger error
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ===== USING CUSTOM ERRORS =====
	fmt.Println("\nCustom Errors:")

	_, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)

		// Type assertion to get more info
		if e, ok := err.(*NegativeSqrtError); ok {
			fmt.Printf("Invalid value was: %f\n", e.Value)
		}
	}

	// ===== ERROR WRAPPING & UNWRAPPING =====
	fmt.Println("\nError Wrapping:")

	err = openFile("non_existent_file.txt")
	if err != nil {
		fmt.Println("Top level error:", err)

		// Unwrap to get the original error
		originalErr := errors.Unwrap(err)
		fmt.Println("Original error:", originalErr)

		// Check if error chain contains specific error
		// errors.Is is preferable to == for wrapped errors
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Confirmed: File does not exist")
		}
	}

	// ===== PANIC AND RECOVER =====
	// Panic is for unrecoverable errors (like array out of bounds)
	// Recover handles panic (similar to catch)

	fmt.Println("\nPanic and Recover:")
	safeFunction()
	fmt.Println("Program continued after panic recovery")
}

func safeFunction() {
	// Defer a function that calls recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("Something went terribly wrong!")
	// Code below panic won't execute
	// fmt.Println("This won't be printed")
}

// Run: go run 13_error_handling.go
