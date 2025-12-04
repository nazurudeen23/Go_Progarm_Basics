// 11_methods.go - Methods in Go

package main

import (
	"fmt"
	"math"
)

// ===== METHODS VS FUNCTIONS =====
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// Method with value receiver
// func (receiver ReceiverType) MethodName(parameters) ReturnType
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with value receiver
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Method with value receiver
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ===== POINTER RECEIVERS =====
// Use pointer receivers when:
// 1. You need to modify the receiver
// 2. The receiver is large (to avoid copying)
// 3. Consistency (if one method needs it, use it for all)

// Method with pointer receiver (can modify the struct)
func (r *Rectangle) Scale(factor float64) {
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}

// Method with value receiver (cannot modify original struct)
func (r Rectangle) TryScale(factor float64) {
	r.Width = r.Width * factor   // Modifies copy only
	r.Height = r.Height * factor // Modifies copy only
}

// ===== METHODS ON NON-STRUCT TYPES =====
// You can define methods on any type defined in the same package
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ===== METHODS AND NIL POINTERS =====
// Methods can be called on nil pointers (unlike some other languages)
type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Sum() int {
	if n == nil {
		return 0
	}
	return n.Value + n.Next.Sum()
}

func main() {
	// ===== BASIC METHODS =====
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	circ := Circle{Radius: 5}
	fmt.Printf("\nCircle Radius: %.2f\n", circ.Radius)
	fmt.Printf("Circle Area: %.2f\n", circ.Area())

	// ===== POINTER VS VALUE RECEIVERS =====
	fmt.Println("\n--- Pointer vs Value Receivers ---")

	// Try to scale with value receiver
	fmt.Println("Before TryScale:", rect)
	rect.TryScale(2)
	fmt.Println("After TryScale:", rect) // Unchanged!

	// Scale with pointer receiver
	fmt.Println("Before Scale:", rect)
	rect.Scale(2)
	fmt.Println("After Scale:", rect) // Changed!

	// Go automatically handles referencing/dereferencing
	// You can call pointer method on value
	rect.Scale(2)

	// You can call value method on pointer
	rectPtr := &rect
	fmt.Printf("Area via pointer: %.2f\n", rectPtr.Area())

	// ===== METHODS ON NON-STRUCTS =====
	f := MyFloat(-math.Sqrt2)
	fmt.Printf("\nMyFloat: %f, Abs: %f\n", f, f.Abs())

	// ===== NIL RECEIVERS =====
	var node *Node // nil pointer
	fmt.Printf("\nSum of nil node: %d\n", node.Sum())

	node = &Node{
		Value: 10,
		Next: &Node{
			Value: 20,
			Next:  nil,
		},
	}
	fmt.Printf("Sum of list: %d\n", node.Sum())
}

// Run: go run 11_methods.go
