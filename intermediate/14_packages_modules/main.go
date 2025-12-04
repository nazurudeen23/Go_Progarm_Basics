// 14_packages_modules.go - Packages and Modules Guide

package main

import (
	"fmt"
	"math/rand" // Import from standard library
	"strings"
)

/*
	===== PACKAGES =====

	1. Every Go file belongs to a package.
	2. The first line of code must be: package <name>
	3. Executable programs must be in package 'main'.
	4. Library code can be in any package name.

	===== VISIBILITY (EXPORTING) =====

	In Go, visibility is controlled by Capitalization.

	- Starts with Capital letter -> Exported (Public)
	  func Calculate() {}
	  type User struct {}
	  var Version string

	- Starts with lowercase letter -> Unexported (Private)
	  func internalHelper() {}
	  type secretConfig struct {}
	  var localCount int

	===== MODULES =====

	A module is a collection of related Go packages.

	Commands:
	1. Initialize a module:
	   go mod init example.com/myproject

	2. Add dependencies (automatic when you build/run):
	   go get github.com/gin-gonic/gin

	3. Clean up dependencies:
	   go mod tidy

	===== DIRECTORY STRUCTURE =====

	myproject/
	├── go.mod
	├── go.sum
	├── main.go           (package main)
	├── utils/
	│   └── string_utils.go (package utils)
	└── models/
	    └── user.go       (package models)

	To use 'utils' in 'main':
	import "example.com/myproject/utils"
*/

func main() {
	// Using standard library packages

	// strings package
	text := "hello go"
	upper := strings.ToUpper(text)
	fmt.Println("Strings package:", upper)

	// math/rand package
	n := rand.Intn(100)
	fmt.Printf("Random number: %d\n", n)

	fmt.Println("\nThis file is a guide. To practice packages:")
	fmt.Println("1. Create a folder 'calculator'")
	fmt.Println("2. Create 'calc.go' inside it with 'package calculator'")
	fmt.Println("3. Add 'func Add(a, b int) int { return a + b }'")
	fmt.Println("4. Import it in your main.go")
}

// Run: go run 14_packages_modules.go
