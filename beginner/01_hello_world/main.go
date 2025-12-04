// 01_hello_world.go - Introduction to Go
// This is your first Go program!

package main // Every Go program starts with a package declaration

import "fmt" // Import the fmt package for formatted I/O

// main() is the entry point of the program
func main() {
	// Print to console
	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go Programming!")
	
	// Single line comment
	/*
		Multi-line comment
		You can write multiple lines here
	*/
	
	// Basic syntax rules:
	// 1. Semicolons are optional (automatically inserted)
	// 2. Code blocks use curly braces {}
	// 3. Opening brace must be on the same line as the statement
	
	fmt.Printf("Go is %s!\n", "awesome")
}

// To run this file: go run 01_hello_world.go
// To build an executable: go build 01_hello_world.go
