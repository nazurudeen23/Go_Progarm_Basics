// 09_structs.go - Structs in Go

package main

import (
	"encoding/json"
	"fmt"
)

// ===== DEFINING STRUCTS =====
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// ===== ANONYMOUS STRUCTS =====
// Can be used for one-off data structures

// ===== EMBEDDED STRUCTS (INHERITANCE-LIKE) =====
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Employee struct {
	Person  // Embedded struct (promoted fields)
	Address // Embedded struct
	Role    string
	Salary  float64
}

// ===== STRUCT TAGS =====
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // Ignore in JSON
	IsActive bool   `json:"is_active,omitempty"`
}

func main() {
	// ===== CREATING STRUCTS =====

	// Method 1: Named fields (Recommended)
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john@example.com",
	}

	// Method 2: Positional (Not recommended, brittle)
	p2 := Person{"Jane", "Smith", 25, "jane@example.com"}

	// Method 3: Zero value
	var p3 Person // All fields are zero values

	// Method 4: Pointer to struct
	p4 := &Person{
		FirstName: "Bob",
		Age:       40,
	}

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
	fmt.Println("Person 3:", p3)
	fmt.Println("Person 4:", p4)

	// ===== ACCESSING FIELDS =====
	fmt.Printf("\nName: %s %s\n", p1.FirstName, p1.LastName)

	// Modifying fields
	p1.Age = 31
	fmt.Printf("New Age: %d\n", p1.Age)

	// Accessing via pointer (Go handles dereferencing automatically)
	p4.LastName = "Builder"
	fmt.Println("Person 4 Updated:", p4)

	// ===== ANONYMOUS STRUCTS =====
	// Useful for temporary data
	config := struct {
		APIKey   string
		Endpoint string
		Timeout  int
	}{
		APIKey:   "xyz123",
		Endpoint: "https://api.example.com",
		Timeout:  5000,
	}

	fmt.Printf("\nConfig: %+v\n", config)

	// ===== EMBEDDED STRUCTS =====
	emp := Employee{
		Person: Person{
			FirstName: "Alice",
			LastName:  "Wonderland",
			Age:       28,
		},
		Address: Address{
			Street: "123 Main St",
			City:   "Tech City",
		},
		Role:   "Developer",
		Salary: 80000,
	}

	// Accessing promoted fields directly
	fmt.Printf("\nEmployee: %s %s\n", emp.FirstName, emp.LastName)
	fmt.Printf("City: %s\n", emp.City)
	fmt.Printf("Role: %s\n", emp.Role)

	// You can still access via the embedded type name if needed
	fmt.Printf("Full Address: %+v\n", emp.Address)

	// ===== STRUCT COMPARISON =====
	// Structs are comparable if all their fields are comparable
	p1Copy := p1

	if p1 == p1Copy {
		fmt.Println("\np1 and p1Copy are equal")
	}

	// ===== CONSTRUCTOR PATTERN =====
	// Go doesn't have constructors, use factory functions
	newPerson := NewPerson("Tom", "Hanks", 60)
	fmt.Println("New Person:", newPerson)

	// ===== STRUCT TAGS AND JSON =====
	user := User{
		ID:       1,
		Username: "gopher",
		Password: "secret_password",
		IsActive: true,
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
	}
	fmt.Printf("\nJSON Output:\n%s\n", string(jsonData))

	// Unmarshal from JSON
	jsonInput := `{"id": 2, "username": "rustacean", "is_active": false}`
	var user2 User
	if err := json.Unmarshal([]byte(jsonInput), &user2); err != nil {
		fmt.Println("Error unmarshaling:", err)
	}
	fmt.Printf("Unmarshaled User: %+v\n", user2)
}

// Factory function
func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Email:     "", // Default value
	}
}

// Run: go run 09_structs.go
