// 06_functions.go - Functions in Go

package main

import "fmt"

// ===== BASIC FUNCTION =====
func greet() {
	fmt.Println("Hello from a function!")
}

// ===== FUNCTION WITH PARAMETERS =====
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// ===== FUNCTION WITH RETURN VALUE =====
func add(a int, b int) int {
	return a + b
}

// Shorthand when parameters have same type
func multiply(a, b int) int {
	return a * b
}

// ===== MULTIPLE RETURN VALUES =====
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// ===== NAMED RETURN VALUES =====
func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Naked return (returns named values)
}

// ===== VARIADIC FUNCTIONS =====
// Accept variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Variadic with other parameters
func printValues(prefix string, values ...int) {
	fmt.Print(prefix)
	for _, v := range values {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// ===== FUNCTIONS AS VALUES =====
func compute(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

// ===== CLOSURES =====
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ===== RECURSIVE FUNCTIONS =====
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// ===== DEFER STATEMENT =====
func demonstrateDefer() {
	defer fmt.Println("This executes last")
	fmt.Println("This executes first")
	fmt.Println("This executes second")
}

func main() {
	// Call basic function
	greet()

	// Function with parameters
	greetPerson("Alice")
	greetPerson("Bob")

	// Function with return value
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	product := multiply(5, 6)
	fmt.Printf("5 x 6 = %d\n", product)

	// Multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	// Division by zero
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Named return values
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle: Area=%.2f, Perimeter=%.2f\n", area, perimeter)

	// Variadic functions
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("Sum of 10,20: %d\n", sum(10, 20))

	// Using slice with variadic function
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of slice: %d\n", sum(numbers...)) // ... unpacks slice

	printValues("Numbers: ", 1, 2, 3, 4, 5)

	// ===== ANONYMOUS FUNCTIONS =====
	// Function without a name
	func() {
		fmt.Println("Anonymous function called!")
	}() // () immediately invokes it

	// Assign to variable
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 5: %d\n", square(5))

	// Functions as parameters
	addResult := compute(add, 10, 5)
	mulResult := compute(multiply, 10, 5)
	fmt.Printf("Compute add: %d\n", addResult)
	fmt.Printf("Compute multiply: %d\n", mulResult)

	// ===== CLOSURES =====
	increment := counter()
	fmt.Println("Counter:", increment()) // 1
	fmt.Println("Counter:", increment()) // 2
	fmt.Println("Counter:", increment()) // 3

	// New closure with independent state
	newCounter := counter()
	fmt.Println("New Counter:", newCounter()) // 1

	// More complex closure
	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	double := multiplier(2)
	triple := multiplier(3)
	fmt.Printf("Double 5: %d\n", double(5)) // 10
	fmt.Printf("Triple 5: %d\n", triple(5)) // 15

	// ===== RECURSIVE FUNCTIONS =====
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Fibonacci sequence (first 10): ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()

	// ===== DEFER =====
	demonstrateDefer()

	// Multiple defers (executed in LIFO order)
	fmt.Println("\nMultiple defers:")
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("Main function")
	// Output order: Main, Third, Second, First
}

// Run: go run 06_functions.go
