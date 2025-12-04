// 23_generics.go - Generics (Go 1.18+)

package main

import "fmt"

// ===== GENERIC FUNCTIONS =====

// Type parameter T can be any type
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Multiple type parameters
func Map[T any, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// ===== TYPE CONSTRAINTS =====

// Constraint: T must be int or float64
func Sum[T int | float64](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Using comparable constraint (types that support == and !=)
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ===== CUSTOM CONSTRAINTS =====

// Define a constraint interface
type Number interface {
	int | int64 | float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// ===== GENERIC TYPES =====

// Generic stack
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Generic pair
type Pair[T, U any] struct {
	First  T
	Second U
}

func main() {
	// ===== GENERIC FUNCTIONS =====
	fmt.Println("=== Generic Print ===")
	Print([]int{1, 2, 3, 4, 5})
	Print([]string{"hello", "world"})

	// ===== MAP FUNCTION =====
	fmt.Println("\n=== Generic Map ===")
	nums := []int{1, 2, 3, 4, 5}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled)

	strings := Map(nums, func(n int) string { return fmt.Sprintf("Number: %d", n) })
	fmt.Println("Strings:", strings)

	// ===== SUM WITH CONSTRAINTS =====
	fmt.Println("\n=== Sum with Constraints ===")
	intSum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("Int sum:", intSum)

	floatSum := Sum([]float64{1.1, 2.2, 3.3})
	fmt.Println("Float sum:", floatSum)

	// ===== CONTAINS =====
	fmt.Println("\n=== Contains ===")
	fmt.Println("Contains 3:", Contains([]int{1, 2, 3, 4}, 3))
	fmt.Println("Contains 'go':", Contains([]string{"hello", "world"}, "go"))

	// ===== MIN WITH CUSTOM CONSTRAINT =====
	fmt.Println("\n=== Min ===")
	fmt.Println("Min(10, 20):", Min(10, 20))
	fmt.Println("Min(3.14, 2.71):", Min(3.14, 2.71))

	// ===== GENERIC STACK =====
	fmt.Println("\n=== Generic Stack ===")

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	for !intStack.IsEmpty() {
		item, _ := intStack.Pop()
		fmt.Println("Popped:", item)
	}

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	item, ok := stringStack.Pop()
	if ok {
		fmt.Println("Popped string:", item)
	}

	// ===== GENERIC PAIR =====
	fmt.Println("\n=== Generic Pair ===")
	p1 := Pair[string, int]{First: "age", Second: 30}
	fmt.Printf("Pair: %s = %d\n", p1.First, p1.Second)

	p2 := Pair[int, string]{First: 1, Second: "Alice"}
	fmt.Printf("Pair: %d = %s\n", p2.First, p2.Second)
}

// Run: go run 23_generics.go
