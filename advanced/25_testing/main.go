// 25_testing.go - Testing in Go

package main

import "fmt"

// ===== TESTING IN GO =====
/*
Go has a built-in testing framework in the 'testing' package.

Test files must:
1. End with _test.go
2. Import "testing"
3. Have functions named TestXxx(t *testing.T)

Example test file (calculator_test.go):

	package main

	import "testing"

	func Add(a, b int) int {
		return a + b
	}

	func TestAdd(t *testing.T) {
		result := Add(2, 3)
		expected := 5

		if result != expected {
			t.Errorf("Add(2, 3) = %d; want %d", result, expected)
		}
	}

	func TestAddNegative(t *testing.T) {
		result := Add(-1, -1)
		expected := -2

		if result != expected {
			t.Errorf("Add(-1, -1) = %d; want %d", result, expected)
		}
	}

Run tests:
	go test
	go test -v          # Verbose
	go test -run TestAdd # Run specific test

===== TABLE-DRIVEN TESTS =====

	func TestAddTable(t *testing.T) {
		tests := []struct {
			name     string
			a, b     int
			expected int
		}{
			{"positive", 2, 3, 5},
			{"negative", -1, -1, -2},
			{"zero", 0, 5, 5},
			{"mixed", -3, 5, 2},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := Add(tt.a, tt.b)
				if result != tt.expected {
					t.Errorf("Add(%d, %d) = %d; want %d",
						tt.a, tt.b, result, tt.expected)
				}
			})
		}
	}

===== HELPER FUNCTIONS =====

	func assertEqual(t *testing.T, got, want int) {
		t.Helper() // Marks this as a helper function
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	func TestWithHelper(t *testing.T) {
		assertEqual(t, Add(2, 3), 5)
	}

===== SETUP AND TEARDOWN =====

	func TestMain(m *testing.M) {
		// Setup
		fmt.Println("Setting up tests...")

		// Run tests
		code := m.Run()

		// Teardown
		fmt.Println("Cleaning up...")

		os.Exit(code)
	}

===== COVERAGE =====

	go test -cover
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

===== EXAMPLE TESTS =====
Example tests serve as documentation and are verified by go test.

	func ExampleAdd() {
		result := Add(2, 3)
		fmt.Println(result)
		// Output: 5
	}
*/

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("This file demonstrates testing concepts.")
	fmt.Println("Create a file named '25_testing_test.go' to write actual tests.")
	fmt.Println("\nExample test:")
	fmt.Println(`
package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}
`)
	fmt.Println("\nRun with: go test -v")
}

// To practice:
// 1. Create 25_testing_test.go
// 2. Write tests for Add and Multiply
// 3. Run: go test -v
