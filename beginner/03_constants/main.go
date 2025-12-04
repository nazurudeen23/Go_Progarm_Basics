// 03_constants.go - Constants in Go

package main

import "fmt"

// Constants can be declared at package level
const Pi = 3.14159
const AppName = "MyApp"

func main() {
	// ===== BASIC CONSTANTS =====
	const greeting = "Hello"
	const maxUsers = 100
	const isEnabled = true

	fmt.Println("Greeting:", greeting)
	fmt.Println("Max Users:", maxUsers)
	fmt.Println("Enabled:", isEnabled)

	// Constants cannot be changed
	// maxUsers = 200 // This would cause a compile error!

	// ===== TYPED CONSTANTS =====
	const typedInt int = 42
	const typedString string = "Go"
	const typedFloat float64 = 3.14

	fmt.Printf("\nTyped Constants:\n")
	fmt.Printf("Int: %d, String: %s, Float: %f\n", typedInt, typedString, typedFloat)

	// ===== UNTYPED CONSTANTS =====
	// Untyped constants have higher precision than typed
	const untypedInt = 42
	const untypedFloat = 3.14159265358979323846

	// Can be assigned to different numeric types
	var i int = untypedInt
	var f float32 = untypedInt
	var f64 float64 = untypedFloat

	fmt.Printf("\nUntyped Constants:\n")
	fmt.Printf("int: %d, float32: %f, float64: %f\n", i, f, f64)

	// ===== MULTIPLE CONSTANTS =====
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)

	fmt.Printf("\nHTTP Status Codes:\n")
	fmt.Printf("OK: %d, Not Found: %d, Error: %d\n", StatusOK, StatusNotFound, StatusError)

	// ===== IOTA - CONSTANT GENERATOR =====
	// iota is used to create incrementing constants
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)

	fmt.Printf("\nDays of Week (iota):\n")
	fmt.Printf("Sunday: %d, Monday: %d, Friday: %d\n", Sunday, Monday, Friday)

	// ===== IOTA WITH EXPRESSIONS =====
	const (
		_  = iota             // 0 (ignored using blank identifier)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
		TB                    // 1 << 40
	)

	fmt.Printf("\nFile Sizes:\n")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	fmt.Printf("1 TB = %d bytes\n", TB)

	// ===== IOTA WITH CUSTOM VALUES =====
	const (
		Red   = iota + 1 // 1
		Green            // 2
		Blue             // 3
	)

	const (
		Apple  = iota * 10 // 0
		Banana             // 10
		Cherry             // 20
	)

	fmt.Printf("\nColors: Red=%d, Green=%d, Blue=%d\n", Red, Green, Blue)
	fmt.Printf("Fruits: Apple=%d, Banana=%d, Cherry=%d\n", Apple, Banana, Cherry)

	// ===== PRACTICAL EXAMPLE: FLAGS =====
	const (
		FlagNone    = 0
		FlagRead    = 1 << iota // 1 << 1 = 2 (binary: 10)
		FlagWrite               // 1 << 2 = 4 (binary: 100)
		FlagExecute             // 1 << 3 = 8 (binary: 1000)
	)

	permissions := FlagRead | FlagWrite // Bitwise OR: 2 | 4 = 6

	fmt.Printf("\nFile Permissions (Bitwise):\n")
	fmt.Printf("Read: %d, Write: %d, Execute: %d\n", FlagRead, FlagWrite, FlagExecute)
	fmt.Printf("Read+Write: %d\n", permissions)
	fmt.Printf("Has Read? %t\n", permissions&FlagRead != 0)
	fmt.Printf("Has Execute? %t\n", permissions&FlagExecute != 0)
}

// Run: go run 03_constants.go
