// 07_arrays_slices.go - Arrays and Slices in Go

package main

import "fmt"

func main() {
	// ===== ARRAYS =====
	// Arrays have fixed size

	// Declaration with size
	var arr1 [5]int // Array of 5 integers, initialized to zero values
	fmt.Println("Empty array:", arr1)

	// Declaration with values
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", arr2)

	// Let compiler count the size
	arr3 := [...]int{10, 20, 30}
	fmt.Println("Auto-sized array:", arr3)

	// Specific indices
	arr4 := [5]int{0: 10, 2: 30, 4: 50}
	fmt.Println("Specific indices:", arr4)

	// Accessing elements
	fmt.Printf("First element: %d\n", arr2[0])
	fmt.Printf("Last element: %d\n", arr2[4])

	// Modifying elements
	arr2[0] = 100
	fmt.Println("Modified array:", arr2)

	// Array length
	fmt.Printf("Array length: %d\n", len(arr2))

	// Iterating over array
	fmt.Println("Array iteration:")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %d\n", i, arr2[i])
	}

	// Using range
	for index, value := range arr2 {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// Multidimensional arrays
	var matrix [3][3]int
	matrix[0][0] = 1
	matrix[1][1] = 5
	matrix[2][2] = 9
	fmt.Println("Matrix:", matrix)

	// ===== SLICES =====
	// Slices are dynamic, flexible views into arrays

	// Creating a slice
	var slice1 []int // nil slice
	fmt.Printf("Nil slice: %v, len=%d, cap=%d, is nil=%t\n",
		slice1, len(slice1), cap(slice1), slice1 == nil)

	// Using make
	slice2 := make([]int, 5) // length 5, capacity 5
	fmt.Printf("Made slice: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))

	slice3 := make([]int, 3, 5) // length 3, capacity 5
	fmt.Printf("Made slice with cap: %v, len=%d, cap=%d\n",
		slice3, len(slice3), cap(slice3))

	// Slice literal
	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice literal:", slice4)

	// Slicing an array
	arr := [5]int{10, 20, 30, 40, 50}
	slice5 := arr[1:4] // Elements at index 1, 2, 3
	fmt.Println("Sliced from array:", slice5)

	// Slicing a slice
	slice6 := slice4[1:3] // Elements at index 1, 2
	fmt.Println("Sliced from slice:", slice6)

	// Omitting bounds
	slice7 := slice4[:3] // From start to index 2
	slice8 := slice4[2:] // From index 2 to end
	slice9 := slice4[:]  // All elements
	fmt.Println("slice[:3]:", slice7)
	fmt.Println("slice[2:]:", slice8)
	fmt.Println("slice[:]:", slice9)

	// ===== APPEND =====
	var numbers []int
	fmt.Println("Initial:", numbers)

	numbers = append(numbers, 1)
	fmt.Println("After append 1:", numbers)

	numbers = append(numbers, 2, 3, 4)
	fmt.Println("After append 2,3,4:", numbers)

	// Append another slice
	moreNumbers := []int{5, 6, 7}
	numbers = append(numbers, moreNumbers...)
	fmt.Println("After append slice:", numbers)

	// ===== CAPACITY AND REALLOCATION =====
	fmt.Println("\nCapacity demonstration:")
	s := make([]int, 0, 3)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	// ===== COPY =====
	source := []int{1, 2, 3}
	dest := make([]int, len(source))
	n := copy(dest, source)
	fmt.Printf("\nCopied %d elements: %v\n", n, dest)

	// Modifying dest doesn't affect source
	dest[0] = 100
	fmt.Println("Source:", source)
	fmt.Println("Dest:", dest)

	// ===== SLICE INTERNALS =====
	// Slices share underlying array
	original := []int{10, 20, 30, 40, 50}
	view := original[1:4]

	fmt.Println("\nSlice internals:")
	fmt.Println("Original:", original)
	fmt.Println("View:", view)

	view[0] = 999 // This modifies the underlying array
	fmt.Println("After modifying view:")
	fmt.Println("Original:", original) // Original changed!
	fmt.Println("View:", view)

	// ===== COMMON SLICE OPERATIONS =====

	// Remove element at index
	nums := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	nums = append(nums[:indexToRemove], nums[indexToRemove+1:]...)
	fmt.Println("\nAfter removing index 2:", nums)

	// Insert element at index
	nums = []int{1, 2, 4, 5}
	indexToInsert := 2
	valueToInsert := 3
	nums = append(nums[:indexToInsert], append([]int{valueToInsert}, nums[indexToInsert:]...)...)
	fmt.Println("After inserting 3 at index 2:", nums)

	// Filter slice
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Println("Even numbers:", evens)

	// ===== 2D SLICES =====
	// Create a 2D slice
	rows, cols := 3, 4
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	// Fill with values
	counter := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = counter
			counter++
		}
	}

	fmt.Println("\n2D Slice (grid):")
	for _, row := range grid {
		fmt.Println(row)
	}

	// ===== PRACTICAL EXAMPLES =====

	// Stack operations
	stack := []int{}
	stack = append(stack, 1, 2, 3) // Push
	fmt.Println("\nStack:", stack)

	top := stack[len(stack)-1]   // Peek
	stack = stack[:len(stack)-1] // Pop
	fmt.Printf("Popped: %d, Stack now: %v\n", top, stack)

	// Queue operations
	queue := []int{}
	queue = append(queue, 1, 2, 3) // Enqueue
	fmt.Println("Queue:", queue)

	front := queue[0] // Peek
	queue = queue[1:] // Dequeue
	fmt.Printf("Dequeued: %d, Queue now: %v\n", front, queue)
}

// Run: go run 07_arrays_slices.go
