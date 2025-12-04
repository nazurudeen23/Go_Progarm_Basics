// 08_maps.go - Maps (Hash Tables) in Go

package main

import "fmt"

func main() {
	// ===== CREATING MAPS =====

	// Declare a nil map
	var map1 map[string]int // nil map, cannot add elements
	fmt.Printf("Nil map: %v, len=%d, is nil=%t\n", map1, len(map1), map1 == nil)

	// Create with make
	map2 := make(map[string]int)
	fmt.Printf("Empty map: %v, len=%d\n", map2, len(map2))

	// Map literal
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}
	fmt.Println("Map literal:", ages)

	// ===== ADDING AND UPDATING =====
	scores := make(map[string]int)

	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Carol"] = 92
	fmt.Println("Scores:", scores)

	// Update value
	scores["Alice"] = 98
	fmt.Println("After update:", scores)

	// ===== ACCESSING VALUES =====
	aliceScore := scores["Alice"]
	fmt.Printf("Alice's score: %d\n", aliceScore)

	// Accessing non-existent key returns zero value
	unknownScore := scores["David"]
	fmt.Printf("David's score: %d\n", unknownScore)

	// Check if key exists
	score, exists := scores["Alice"]
	if exists {
		fmt.Printf("Alice's score exists: %d\n", score)
	}

	score, exists = scores["David"]
	if !exists {
		fmt.Println("David's score doesn't exist")
	}

	// ===== DELETING ELEMENTS =====
	fmt.Println("\nBefore delete:", scores)
	delete(scores, "Bob")
	fmt.Println("After delete Bob:", scores)

	// Deleting non-existent key is safe
	delete(scores, "NonExistent")

	// ===== ITERATING OVER MAPS =====
	fmt.Println("\nIterating over map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Only keys
	fmt.Println("\nOnly keys:")
	for name := range ages {
		fmt.Println(name)
	}

	// Note: Map iteration order is random!
	fmt.Println("\nIteration order is random:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d: ", i+1)
		for name := range ages {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}

	// ===== MAP LENGTH =====
	fmt.Printf("\nMap length: %d\n", len(ages))

	// ===== MAPS WITH DIFFERENT TYPES =====

	// Map with struct values
	type Person struct {
		Name string
		Age  int
	}

	people := map[int]Person{
		1: {Name: "Alice", Age: 25},
		2: {Name: "Bob", Age: 30},
		3: {Name: "Carol", Age: 35},
	}
	fmt.Println("\nMap with struct values:", people)

	// Map with slice key is NOT allowed (slices are not comparable)
	// But map with array key is allowed
	arrayKeyMap := map[[2]int]string{
		{0, 0}: "origin",
		{1, 2}: "point A",
	}
	fmt.Println("Map with array keys:", arrayKeyMap)

	// Nested maps
	users := map[string]map[string]string{
		"user1": {
			"name":  "Alice",
			"email": "alice@example.com",
		},
		"user2": {
			"name":  "Bob",
			"email": "bob@example.com",
		},
	}
	fmt.Println("\nNested maps:", users)
	fmt.Printf("User1 email: %s\n", users["user1"]["email"])

	// ===== MAP AS SET =====
	// Use map[T]bool or map[T]struct{} for sets

	set := make(map[string]bool)
	set["apple"] = true
	set["banana"] = true
	set["cherry"] = true

	// Check membership
	if set["apple"] {
		fmt.Println("\nApple is in the set")
	}

	if !set["orange"] {
		fmt.Println("Orange is not in the set")
	}

	// Using struct{} is more memory efficient
	efficientSet := make(map[string]struct{})
	efficientSet["go"] = struct{}{}
	efficientSet["rust"] = struct{}{}

	if _, exists := efficientSet["go"]; exists {
		fmt.Println("Go is in the efficient set")
	}

	// ===== PRACTICAL EXAMPLES =====

	// Example 1: Count word frequencies
	// text := "hello world hello go go go" // Unused
	words := []string{"hello", "world", "hello", "go", "go", "go"}

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("\nWord frequencies:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d time(s)\n", word, count)
	}

	// Example 2: Group by category
	items := []struct {
		category string
		name     string
	}{
		{"fruit", "apple"},
		{"fruit", "banana"},
		{"vegetable", "carrot"},
		{"fruit", "cherry"},
		{"vegetable", "broccoli"},
	}

	grouped := make(map[string][]string)
	for _, item := range items {
		grouped[item.category] = append(grouped[item.category], item.name)
	}

	fmt.Println("\nGrouped items:")
	for category, names := range grouped {
		fmt.Printf("%s: %v\n", category, names)
	}

	// Example 3: Memoization (caching)
	cache := make(map[int]int)

	fibonacci := func(n int) int {
		if n <= 1 {
			return n
		}

		// Check cache
		if val, exists := cache[n]; exists {
			return val
		}

		// Compute and store in cache
		// (This is simplified, real memoized fibonacci needs recursion with cache)
		result := n // Simplified
		cache[n] = result
		return result
	}

	fmt.Println("\nFibonacci with cache:")
	for i := 0; i < 5; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fibonacci(i))
	}
	fmt.Println("Cache:", cache)

	// Example 4: Two-way lookup
	idToName := map[int]string{
		1: "Alice",
		2: "Bob",
		3: "Carol",
	}

	nameToID := make(map[string]int)
	for id, name := range idToName {
		nameToID[name] = id
	}

	fmt.Println("\nTwo-way lookup:")
	fmt.Printf("ID 2 -> %s\n", idToName[2])
	fmt.Printf("Bob -> ID %d\n", nameToID["Bob"])

	// ===== COPYING MAPS =====
	// Maps are reference types, assignment doesn't copy
	original := map[string]int{"a": 1, "b": 2}
	reference := original // Same underlying map
	reference["a"] = 100

	fmt.Println("\nMaps are reference types:")
	fmt.Println("Original:", original) // Modified!
	fmt.Println("Reference:", reference)

	// To copy, must iterate
	original = map[string]int{"a": 1, "b": 2}
	copied := make(map[string]int)
	for k, v := range original {
		copied[k] = v
	}
	copied["a"] = 100

	fmt.Println("\nProper copy:")
	fmt.Println("Original:", original) // Not modified
	fmt.Println("Copied:", copied)
}

// Run: go run 08_maps.go
