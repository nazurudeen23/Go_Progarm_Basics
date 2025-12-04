// 12_interfaces.go - Interfaces in Go

package main

import (
	"fmt"
	"math"
)

// ===== DEFINING INTERFACES =====
// An interface is a set of method signatures
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ===== IMPLEMENTING INTERFACES =====
// A type implements an interface by implementing its methods.
// There is no "implements" keyword (implicit implementation).

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// ===== INTERFACE AS PARAMETER =====
func printShapeInfo(s Shape) {
	fmt.Printf("Shape Type: %T\n", s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println("---")
}

// ===== EMPTY INTERFACE =====
// interface{} has no methods, so all types implement it.
// Used to handle values of unknown type.
// (In Go 1.18+, 'any' is an alias for interface{})
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// ===== TYPE ASSERTION =====
// Accessing the underlying concrete value
func checkType(i interface{}) {
	// Unsafe assertion (panics if wrong type)
	// s := i.(string)

	// Safe assertion
	s, ok := i.(string)
	if ok {
		fmt.Printf("It's a string: %s\n", s)
	} else {
		fmt.Printf("Not a string\n")
	}
}

// ===== TYPE SWITCH =====
func doTypeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("It is boolean: %t\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// ===== INTERFACE COMPOSITION =====
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

// ReadWriter embeds Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	// ===== USING INTERFACES =====
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 5}

	// Both satisfy Shape interface
	var s Shape

	s = r
	fmt.Println("Rectangle assigned to Shape:")
	printShapeInfo(s)

	s = c
	fmt.Println("Circle assigned to Shape:")
	printShapeInfo(s)

	// Slice of interfaces
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2},
		Rectangle{Width: 5, Height: 5},
	}

	fmt.Println("Iterating over shapes:")
	for _, shape := range shapes {
		printShapeInfo(shape)
	}

	// ===== EMPTY INTERFACE =====
	fmt.Println("\nEmpty Interface:")
	describe(42)
	describe("hello")
	describe(true)
	describe(r)

	// ===== TYPE ASSERTION =====
	fmt.Println("\nType Assertion:")
	var x interface{} = "hello"
	checkType(x)

	x = 123
	checkType(x)

	// ===== TYPE SWITCH =====
	fmt.Println("\nType Switch:")
	doTypeSwitch(21)
	doTypeSwitch("hello")
	doTypeSwitch(true)
	doTypeSwitch(3.14)
}

// Run: go run 12_interfaces.go
