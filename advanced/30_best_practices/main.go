// 30_best_practices.go - Go Best Practices and Idioms

package main

import (
	"errors"
	"fmt"
	"time"
)

// ===== GO BEST PRACTICES =====

/*
1. CODE ORGANIZATION
   - One package per directory
   - Package names should be short, lowercase, no underscores
   - Exported names start with capital letter
   - Use meaningful names (avoid single letters except for short scopes)

2. ERROR HANDLING
   - Always check errors
   - Don't use panic for normal error handling
   - Return errors, don't log and return
   - Wrap errors with context: fmt.Errorf("context: %w", err)

3. CONCURRENCY
   - Don't communicate by sharing memory; share memory by communicating
   - Use channels to pass ownership
   - Use sync.WaitGroup to wait for goroutines
   - Always close channels when done sending

4. INTERFACES
   - Accept interfaces, return structs
   - Keep interfaces small (1-3 methods)
   - Define interfaces where they're used, not where they're implemented

5. NAMING CONVENTIONS
   - Use camelCase, not snake_case
   - Acronyms should be all caps: HTTP, URL, ID
   - Getters don't use "Get" prefix: obj.Name(), not obj.GetName()
   - Setters use "Set" prefix: obj.SetName()

6. COMMENTS
   - Package comment on package declaration
   - Exported items should have doc comments
   - Comments should be complete sentences
   - Start with the name of the item
*/

// ===== EXAMPLES =====

// Good: Clear, descriptive name
type UserRepository struct {
	db Database
}

// Bad: Unclear abbreviation
// type UsrRepo struct {}

// Good: Interface where it's used
type Database interface {
	Query(sql string) ([]Row, error)
}

type Row struct{}

// Good: Small, focused interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Good: Error handling
func ProcessData(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty data")
	}

	// Process...

	return nil
}

// Bad: Panic for normal errors
// func ProcessDataBad(data []byte) {
// 	if len(data) == 0 {
// 		panic("empty data")
// 	}
// }

// Good: Return error with context
func LoadConfig(filename string) (*Config, error) {
	data, err := readFile(filename)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	config, err := parseConfig(data)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return config, nil
}

type Config struct{}

func readFile(filename string) ([]byte, error) {
	return nil, errors.New("file not found")
}

func parseConfig(data []byte) (*Config, error) {
	return &Config{}, nil
}

// Good: Constructor pattern
func NewUserRepository(db Database) *UserRepository {
	return &UserRepository{db: db}
}

// Good: Options pattern for many parameters
type ServerOptions struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(opts ServerOptions) *Server {
	return &Server{
		port:         opts.Port,
		readTimeout:  opts.ReadTimeout,
		writeTimeout: opts.WriteTimeout,
	}
}

type Server struct {
	port         int
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// Good: Functional options pattern (advanced)
type ServerOption func(*Server)

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.readTimeout = timeout
		s.writeTimeout = timeout
	}
}

func NewServerWithOptions(opts ...ServerOption) *Server {
	s := &Server{
		port:         8080, // defaults
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Good: Table-driven tests (see 25_testing.go)

// Good: Use defer for cleanup
func ProcessFile(filename string) error {
	// f, err := os.Open(filename)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close() // Always close

	// Process file...

	return nil
}

// Good: Zero values are useful
type Counter struct {
	count int // Starts at 0, which is useful
}

func (c *Counter) Increment() {
	c.count++
}

// Good: Make the zero value useful
type Buffer struct {
	data []byte // nil slice is valid
}

func (b *Buffer) Write(p []byte) {
	b.data = append(b.data, p...) // Works even if data is nil
}

func main() {
	fmt.Println("=== Go Best Practices ===\n")

	// Using functional options
	server := NewServerWithOptions(
		WithPort(9000),
		WithTimeout(10*time.Second),
	)
	fmt.Printf("Server: %+v\n", server)

	// Zero value usage
	var counter Counter
	counter.Increment()
	counter.Increment()
	fmt.Printf("Counter: %d\n", counter.count)

	var buf Buffer
	buf.Write([]byte("hello"))
	fmt.Printf("Buffer: %s\n", buf.data)
}

/*
===== MORE BEST PRACTICES =====

1. PERFORMANCE
   - Use benchmarks to measure, don't guess
   - Preallocate slices when size is known: make([]int, 0, expectedSize)
   - Use sync.Pool for frequently allocated objects
   - Avoid premature optimization

2. TESTING
   - Write table-driven tests
   - Use t.Helper() in test helpers
   - Test exported API, not implementation details
   - Use testdata/ directory for test fixtures

3. DEPENDENCIES
   - Use go.mod for dependency management
   - Keep dependencies minimal
   - Vendor dependencies for critical projects

4. DOCUMENTATION
   - Use godoc format
   - Provide examples in _test.go files
   - Keep README.md updated

5. CODE STYLE
   - Run gofmt (or use goimports)
   - Use golint and go vet
   - Follow Effective Go: https://go.dev/doc/effective_go

6. PROJECT STRUCTURE
   cmd/          - Main applications
   pkg/          - Library code
   internal/     - Private code
   api/          - API definitions
   web/          - Web assets
   scripts/      - Build scripts
   test/         - Additional test files

===== USEFUL TOOLS =====

- gofmt: Format code
- goimports: Manage imports
- golint: Linting
- go vet: Static analysis
- staticcheck: Advanced static analysis
- gopls: Language server
- delve: Debugger
*/

// This file is a guide - read through the comments!
