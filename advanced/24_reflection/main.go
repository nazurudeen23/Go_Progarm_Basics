// 24_reflection.go - Reflection in Go

package main

import (
	"fmt"
	"reflect"
)

// ===== REFLECTION =====
// Reflection allows inspection of types and values at runtime.
// Use sparingly - it's powerful but can be slow and error-prone.

type Person struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"min=0,max=120"`
	Email string `json:"email"`
}

func main() {
	// ===== TYPE AND VALUE =====
	var x float64 = 3.4

	fmt.Println("Type:", reflect.TypeOf(x))
	fmt.Println("Value:", reflect.ValueOf(x))

	// ===== KIND =====
	// Kind is the specific kind of type (int, struct, slice, etc.)
	v := reflect.ValueOf(x)
	fmt.Println("Kind:", v.Kind())

	// ===== INSPECTING STRUCTS =====
	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

	t := reflect.TypeOf(p)
	fmt.Printf("\nType: %v\n", t)
	fmt.Printf("Kind: %v\n", t.Kind())
	fmt.Printf("Number of fields: %d\n", t.NumField())

	// Iterate over fields
	fmt.Println("\nStruct Fields:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("  Field %d: %s (type: %v, tag: %s)\n",
			i, field.Name, field.Type, field.Tag)
	}

	// ===== GETTING FIELD VALUES =====
	val := reflect.ValueOf(p)
	fmt.Println("\nField Values:")
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := t.Field(i)
		fmt.Printf("  %s = %v\n", fieldType.Name, fieldValue.Interface())
	}

	// ===== STRUCT TAGS =====
	fmt.Println("\nStruct Tags:")
	nameField, _ := t.FieldByName("Name")
	jsonTag := nameField.Tag.Get("json")
	validateTag := nameField.Tag.Get("validate")
	fmt.Printf("Name field - json: %s, validate: %s\n", jsonTag, validateTag)

	// ===== MODIFYING VALUES =====
	// Can only modify if value is addressable (pointer)
	fmt.Println("\n=== Modifying Values ===")

	p2 := Person{Name: "Bob", Age: 25}
	fmt.Println("Before:", p2)

	// Get pointer value
	ptrVal := reflect.ValueOf(&p2)
	// Get the element the pointer points to
	elemVal := ptrVal.Elem()

	// Modify field
	nameField2 := elemVal.FieldByName("Name")
	if nameField2.CanSet() {
		nameField2.SetString("Robert")
	}

	ageField := elemVal.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(26)
	}

	fmt.Println("After:", p2)

	// ===== CALLING METHODS =====
	fmt.Println("\n=== Calling Methods ===")

	s := "hello"
	val = reflect.ValueOf(s)

	// Get method by name
	method := val.MethodByName("ToUpper")
	if method.IsValid() {
		// Note: strings don't have ToUpper method, this is just an example
		fmt.Println("Method found")
	}

	// ===== TYPE SWITCH VS REFLECTION =====
	fmt.Println("\n=== Inspecting Interface{} ===")

	var i interface{} = 42
	inspectValue(i)

	i = "hello"
	inspectValue(i)

	i = []int{1, 2, 3}
	inspectValue(i)
}

func inspectValue(x interface{}) {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)

	fmt.Printf("Value: %v, Type: %v, Kind: %v\n", v, t, v.Kind())

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("  It's an integer: %d\n", v.Int())
	case reflect.String:
		fmt.Printf("  It's a string: %s\n", v.String())
	case reflect.Slice:
		fmt.Printf("  It's a slice with %d elements\n", v.Len())
	}
}

// Run: go run 24_reflection.go
