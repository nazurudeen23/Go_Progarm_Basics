// 10_pointers.go - Pointers in Go

package main

import "fmt"

func main() {
	// ===== BASIC POINTERS =====
	// A pointer holds the memory address of a value

	x := 10
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x) // & operator generates a pointer

	// Declaring a pointer
	var p *int
	p = &x

	fmt.Printf("Value of p (address of x): %p\n", p)
	fmt.Printf("Value through p (dereferencing): %d\n", *p) // * operator dereferences

	// Changing value through pointer
	*p = 20
	fmt.Printf("New value of x: %d\n", x)

	// ===== POINTER TO POINTER =====
	var pp **int
	pp = &p

	fmt.Printf("\nPointer to pointer (pp): %p\n", pp)
	fmt.Printf("Dereferencing pp gives p: %p\n", *pp)
	fmt.Printf("Double dereferencing gives x: %d\n", **pp)

	// ===== PASSING BY VALUE VS REFERENCE =====

	// Pass by value
	num := 5
	modifyValue(num)
	fmt.Printf("\nAfter modifyValue: %d (Unchanged)\n", num)

	// Pass by reference (pointer)
	modifyPointer(&num)
	fmt.Printf("After modifyPointer: %d (Changed)\n", num)

	// ===== NIL POINTERS =====
	var nilPtr *int
	fmt.Printf("\nNil pointer: %v\n", nilPtr)

	if nilPtr == nil {
		fmt.Println("Pointer is nil")
	}

	// DANGER: Dereferencing nil pointer causes panic
	// fmt.Println(*nilPtr) // Runtime error!

	// ===== NEW FUNCTION =====
	// new(T) allocates zeroed storage for T and returns *T
	ptr := new(int)
	fmt.Printf("\nCreated with new(): %d\n", *ptr)
	*ptr = 100
	fmt.Printf("Updated value: %d\n", *ptr)

	// ===== POINTERS TO STRUCTS =====
	// type Config struct {
	// 	Port    int
	// 	Enabled bool
	// }

	// Struct pointer
	conf := &Config{Port: 8080, Enabled: true}

	// Accessing fields (auto-dereference)
	fmt.Printf("\nStruct pointer port: %d\n", conf.Port)
	// Equivalent to:
	fmt.Printf("Explicit dereference: %d\n", (*conf).Port)

	// Modifying struct via function
	updateConfig(conf)
	fmt.Printf("Updated config: %+v\n", conf)

	// ===== POINTER RECEIVERS (METHODS) =====
	// We'll see this more in methods section, but here's a preview

	// ===== PERFORMANCE IMPLICATIONS =====
	// Use pointers when:
	// 1. You need to modify the original value
	// 2. The struct is large and copying is expensive
	// 3. You want to signify "optional" values (nil = missing)

	// Don't use pointers when:
	// 1. The type is small (int, bool, small struct)
	// 2. You want to ensure immutability
	// 3. Slices, maps, channels (they are already reference-like)

	// ===== SLICES AND MAPS ARE REFERENCE TYPES =====
	// They contain pointers internally, so passing them to functions
	// allows modification of underlying data without explicit pointers

	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Printf("\nSlice after function: %v (Modified!)\n", slice)
}

func modifyValue(n int) {
	n = 100 // Only modifies local copy
}

func modifyPointer(n *int) {
	*n = 100 // Modifies value at address
}

type Config struct {
	Port    int
	Enabled bool
}

func updateConfig(c *Config) {
	c.Port = 9090
	c.Enabled = false
}

func modifySlice(s []int) {
	s[0] = 999
}

// Run: go run 10_pointers.go
