// 20_json.go - JSON Handling in Go

package main

import (
	"encoding/json"
	"fmt"
)

// Struct tags control JSON encoding/decoding
type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Tags        []string `json:"tags,omitempty"` // Omit if empty
	IsAvailable bool     `json:"is_available"`
	private     string   // Unexported fields are ignored
}

func main() {
	// ===== ENCODING (MARSHALING) =====

	// 1. Struct to JSON
	p1 := Product{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		Tags:        []string{"electronics", "computer"},
		IsAvailable: true,
		private:     "secret",
	}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("JSON Output:", string(jsonData))

	// Pretty print
	prettyJSON, _ := json.MarshalIndent(p1, "", "  ")
	fmt.Println("\nPretty JSON:\n", string(prettyJSON))

	// 2. Map to JSON
	values := map[string]interface{}{
		"user":   "alice",
		"age":    30,
		"active": true,
	}
	mapJSON, _ := json.Marshal(values)
	fmt.Println("\nMap to JSON:", string(mapJSON))

	// 3. Slice to JSON
	slice := []string{"apple", "banana", "cherry"}
	sliceJSON, _ := json.Marshal(slice)
	fmt.Println("Slice to JSON:", string(sliceJSON))

	// ===== DECODING (UNMARSHALING) =====

	// 1. JSON to Struct
	jsonStr := `{"id": 2, "name": "Phone", "price": 599.50, "is_available": true}`
	var p2 Product

	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	fmt.Printf("\nDecoded Struct: %+v\n", p2)

	// 2. JSON to Map (Generic)
	// Useful when structure is unknown
	jsonStr2 := `{"result": "success", "data": [1, 2, 3]}`
	var result map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr2), &result); err != nil {
		panic(err)
	}
	fmt.Println("Decoded Map:", result)

	// Accessing generic map values requires type assertion
	status := result["result"].(string)
	fmt.Println("Status:", status)

	data := result["data"].([]interface{})
	fmt.Println("First data item:", data[0])

	// ===== CUSTOM JSON ENCODING =====
	// You can implement Marshaler/Unmarshaler interfaces
}

// Run: go run 20_json.go
