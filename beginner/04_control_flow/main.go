// 04_control_flow.go - If/Else and Switch Statements

package main

import (
	"fmt"
	"time"
)

func main() {
	// ===== IF STATEMENT =====
	age := 20

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// ===== IF-ELSE =====
	temperature := 25

	if temperature > 30 {
		fmt.Println("It's hot!")
	} else {
		fmt.Println("It's comfortable")
	}

	// ===== IF-ELSE IF-ELSE =====
	score := 75

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// ===== IF WITH INITIALIZATION =====
	// You can initialize a variable in the if statement
	if num := 10; num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
	// Note: 'num' is only accessible within if-else block

	// ===== SWITCH STATEMENT =====
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday": // Multiple values
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek day")
	}

	// ===== SWITCH WITH EXPRESSIONS =====
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	case hour < 21:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night!")
	}

	// ===== SWITCH WITH INITIALIZATION =====
	switch month := time.Now().Month(); month {
	case time.January, time.February, time.December:
		fmt.Printf("%s is in winter\n", month)
	case time.March, time.April, time.May:
		fmt.Printf("%s is in spring\n", month)
	case time.June, time.July, time.August:
		fmt.Printf("%s is in summer\n", month)
	default:
		fmt.Printf("%s is in fall\n", month)
	}

	// ===== SWITCH WITH FALLTHROUGH =====
	// By default, Go doesn't fall through cases
	// Use 'fallthrough' keyword to continue to next case
	num := 2

	fmt.Printf("\nSwitch with fallthrough for %d:\n", num)
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough // This will execute the next case
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	}

	// ===== TYPE SWITCH =====
	// Used with interfaces (we'll learn more about this later)
	var value interface{} = "hello"

	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// ===== LOGICAL OPERATORS =====
	x := 10
	y := 20

	// AND (&&)
	if x > 5 && y > 15 {
		fmt.Println("Both conditions are true")
	}

	// OR (||)
	if x < 5 || y > 15 {
		fmt.Println("At least one condition is true")
	}

	// NOT (!)
	isRaining := false
	if !isRaining {
		fmt.Println("It's not raining")
	}

	// ===== COMPARISON OPERATORS =====
	a, b := 10, 20

	fmt.Printf("\nComparison Operators (a=%d, b=%d):\n", a, b)
	fmt.Printf("a == b: %t\n", a == b) // Equal
	fmt.Printf("a != b: %t\n", a != b) // Not equal
	fmt.Printf("a < b: %t\n", a < b)   // Less than
	fmt.Printf("a <= b: %t\n", a <= b) // Less than or equal
	fmt.Printf("a > b: %t\n", a > b)   // Greater than
	fmt.Printf("a >= b: %t\n", a >= b) // Greater than or equal
}

// Run: go run 04_control_flow.go
