// 19_file_io.go - File I/O in Go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"

	// ===== WRITING TO FILES =====

	// Method 1: os.WriteFile (Quick, for smaller files)
	content := []byte("Hello, Go!\nThis is a test file.\n")
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File written successfully")

	// Method 2: os.Create + Write (More control)
	f, err := os.Create("lines.txt")
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(w, "Line %d\n", i)
	}
	w.Flush() // Ensure all data is written
	f.Close()
	fmt.Println("Lines written successfully")

	// ===== READING FROM FILES =====

	// Method 1: os.ReadFile (Read entire file)
	fmt.Println("\n--- Reading entire file ---")
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))

	// Method 2: Open + Read (Chunk by chunk)
	fmt.Println("\n--- Reading in chunks ---")
	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek to known location
	o2, err := f.Seek(6, 0)
	if err != nil {
		panic(err)
	}
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2]))

	// Method 3: bufio.Scanner (Line by line)
	fmt.Println("\n--- Reading line by line ---")

	f2, err := os.Open("lines.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	scanner := bufio.NewScanner(f2)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}

	// ===== CLEANUP =====
	os.Remove(filename)
	os.Remove("lines.txt")

	// ===== FILE INFO =====
	fileInfo, err := os.Stat("19_file_io.go")
	if err == nil {
		fmt.Println("\n--- File Info ---")
		fmt.Println("Name:", fileInfo.Name())
		fmt.Println("Size:", fileInfo.Size(), "bytes")
		fmt.Println("Permissions:", fileInfo.Mode())
		fmt.Println("Last Modified:", fileInfo.ModTime())
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Run: go run 19_file_io.go
