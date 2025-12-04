// 02_variables.go - Variables and Data Types

package main

import "fmt"

func main() {
	// ===== VARIABLE DECLARATION =====

	// Method 1: var keyword with type
	var name string = "John"
	var age int = 30
	var isStudent bool = true

	// Method 2: var keyword with type inference
	var city = "New York" // Go infers this is a string
	var score = 95.5      // Go infers this is float64

	// Method 3: Short declaration (inside functions only)
	country := "USA" // := declares and assigns
	height := 5.9

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("City:", city)
	fmt.Println("Score:", score)
	fmt.Println("Country:", country)
	fmt.Println("Height:", height)

	// ===== BASIC DATA TYPES =====

	// Integers
	var smallInt int8 = 127           // -128 to 127
	var mediumInt int16 = 32767       // -32768 to 32767
	var regularInt int32 = 2147483647 // -2^31 to 2^31-1
	var bigInt int64 = 9223372036854775807

	// Unsigned integers
	var uSmall uint8 = 255 // 0 to 255
	var uMedium uint16 = 65535

	// Float
	var pi float32 = 3.14159
	var precisePi float64 = 3.14159265359

	// Complex numbers
	var complexNum complex64 = 1 + 2i

	// Byte (alias for uint8)
	var b byte = 'A'

	// Rune (alias for int32, represents a Unicode code point)
	var r rune = '世'

	fmt.Printf("\nData Types:\n")
	fmt.Printf("int8: %d\n", smallInt)
	fmt.Printf("int16: %d\n", mediumInt)
	fmt.Printf("int32: %d\n", regularInt)
	fmt.Printf("int64: %d\n", bigInt)
	fmt.Printf("uint8: %d\n", uSmall)
	fmt.Printf("uint16: %d\n", uMedium)
	fmt.Printf("float32: %f\n", pi)
	fmt.Printf("float64: %f\n", precisePi)
	fmt.Printf("complex64: %v\n", complexNum)
	fmt.Printf("byte: %c (value: %d)\n", b, b)
	fmt.Printf("rune: %c (value: %d)\n", r, r)

	// ===== ZERO VALUES =====
	// Variables declared without initialization get zero values
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultBool bool     // false
	var defaultString string // "" (empty string)

	fmt.Printf("\nZero Values:\n")
	fmt.Printf("int: %d\n", defaultInt)
	fmt.Printf("float64: %f\n", defaultFloat)
	fmt.Printf("bool: %t\n", defaultBool)
	fmt.Printf("string: '%s'\n", defaultString)

	// ===== TYPE CONVERSION =====
	var x int = 42
	var y float64 = float64(x) // Must be explicit
	var z uint = uint(y)

	fmt.Printf("\nType Conversion:\n")
	fmt.Printf("int: %d -> float64: %f -> uint: %d\n", x, y, z)

	// ===== MULTIPLE VARIABLE DECLARATION =====
	var (
		firstName string = "Alice"
		lastName  string = "Smith"
		userAge   int    = 25
	)

	// Multiple assignment
	a, b, c := 1, 2, 3

	fmt.Printf("\nMultiple Variables:\n")
	fmt.Printf("Name: %s %s, Age: %d\n", firstName, lastName, userAge)
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}

// Run: go run 02_variables.go
