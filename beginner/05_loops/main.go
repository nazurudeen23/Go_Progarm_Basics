// 05_loops.go - Loops in Go

package main

import "fmt"

func main() {
	// ===== BASIC FOR LOOP =====
	// Go has only one looping construct: the for loop

	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// ===== WHILE-STYLE LOOP =====
	// Go doesn't have 'while', but for can work like while
	fmt.Println("\nWhile-style loop:")
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// ===== INFINITE LOOP =====
	// Can be broken with 'break'
	fmt.Println("\nInfinite loop with break:")
	n := 0
	for {
		if n >= 3 {
			break // Exit the loop
		}
		fmt.Printf("n = %d\n", n)
		n++
	}

	// ===== CONTINUE STATEMENT =====
	// Skip to next iteration
	fmt.Println("\nLoop with continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip when i is 2
		}
		fmt.Printf("Value: %d\n", i)
	}

	// ===== RANGE OVER ARRAYS/SLICES =====
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("\nRange over slice:")
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// If you only need the value
	fmt.Println("\nRange with only values:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// If you only need the index
	fmt.Println("\nRange with only indices:")
	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
	}

	// ===== RANGE OVER MAPS =====
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}

	fmt.Println("\nRange over map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// ===== RANGE OVER STRINGS =====
	// Iterates over Unicode code points (runes)
	text := "Hello, 世界"

	fmt.Println("\nRange over string (runes):")
	for index, char := range text {
		fmt.Printf("Index %d: %c (Unicode: %U)\n", index, char, char)
	}

	// ===== NESTED LOOPS =====
	fmt.Println("\nNested loops - Multiplication table:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// ===== LABELED BREAKS =====
	// Break out of outer loop from inner loop
	fmt.Println("\nLabeled break:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 1 && j == 1 {
				break outer // Break out of the outer loop
			}
		}
	}

	// ===== LABELED CONTINUE =====
	fmt.Println("\nLabeled continue:")
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outerLoop // Continue the outer loop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// ===== PRACTICAL EXAMPLES =====

	// Example 1: Find first even number
	nums := []int{1, 3, 5, 8, 9, 10}
	fmt.Println("\nFind first even number:")
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("First even number: %d\n", num)
			break
		}
	}

	// Example 2: Sum of numbers
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Sum: %d\n", total)

	// Example 3: FizzBuzz
	fmt.Println("\nFizzBuzz (1-15):")
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	// Example 4: Reverse iteration
	fmt.Println("\nReverse iteration:")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Printf("%d ", numbers[i])
	}
	fmt.Println()
}

// Run: go run 05_loops.go
