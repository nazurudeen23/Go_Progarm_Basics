// 29_database.go - Database Operations

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// ===== DATABASE OPERATIONS =====
/*
This example uses SQLite, but the same patterns apply to PostgreSQL, MySQL, etc.

Install SQLite driver:
	go get github.com/mattn/go-sqlite3

For PostgreSQL:
	go get github.com/lib/pq

For MySQL:
	go get github.com/go-sql-driver/mysql

Connection strings:
	SQLite:     "file:test.db?cache=shared&mode=memory"
	PostgreSQL: "postgres://user:password@localhost/dbname?sslmode=disable"
	MySQL:      "user:password@tcp(localhost:3306)/dbname"
*/

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	// ===== CONNECT TO DATABASE =====
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")

	// ===== CREATE TABLE =====
	createTableSQL := `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created!")

	// ===== INSERT DATA =====
	insertSQL := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"

	result, err := db.Exec(insertSQL, "Alice", "alice@example.com", 30)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := result.LastInsertId()
	fmt.Printf("Inserted user with ID: %d\n", id)

	// Insert multiple
	users := []User{
		{Name: "Bob", Email: "bob@example.com", Age: 25},
		{Name: "Carol", Email: "carol@example.com", Age: 35},
	}

	for _, user := range users {
		db.Exec(insertSQL, user.Name, user.Email, user.Age)
	}

	// ===== QUERY SINGLE ROW =====
	var user User
	querySQL := "SELECT id, name, email, age FROM users WHERE id = ?"

	err = db.QueryRow(querySQL, 1).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("\nUser: %+v\n", user)

	// ===== QUERY MULTIPLE ROWS =====
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nAll users:")
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  %+v\n", u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// ===== UPDATE =====
	updateSQL := "UPDATE users SET age = ? WHERE name = ?"
	result, err = db.Exec(updateSQL, 31, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nUpdated %d row(s)\n", rowsAffected)

	// ===== DELETE =====
	deleteSQL := "DELETE FROM users WHERE name = ?"
	result, err = db.Exec(deleteSQL, "Bob")
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ = result.RowsAffected()
	fmt.Printf("Deleted %d row(s)\n", rowsAffected)

	// ===== TRANSACTIONS =====
	fmt.Println("\nTransaction example:")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(insertSQL, "David", "david@example.com", 40)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(insertSQL, "Eve", "eve@example.com", 28)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction committed!")

	// ===== PREPARED STATEMENTS =====
	// More efficient for repeated queries
	stmt, err := db.Prepare("SELECT name, age FROM users WHERE age > ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err = stmt.Query(30)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nUsers older than 30:")
	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age)
		fmt.Printf("  %s (%d)\n", name, age)
	}
}

/*
===== BEST PRACTICES =====

1. Always close resources:
   defer db.Close()
   defer rows.Close()
   defer stmt.Close()

2. Use prepared statements for repeated queries

3. Use transactions for multiple related operations

4. Handle sql.ErrNoRows separately from other errors

5. Use connection pooling (built-in):
   db.SetMaxOpenConns(25)
   db.SetMaxIdleConns(5)
   db.SetConnMaxLifetime(5 * time.Minute)

6. Use context for timeouts:
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   rows, err := db.QueryContext(ctx, query)

===== POPULAR ORMs =====

1. GORM: github.com/go-gorm/gorm
2. sqlx: github.com/jmoiron/sqlx (not an ORM, but helpful extensions)
3. ent: entgo.io
*/

// Run: go run 29_database.go
// Note: You need to install the SQLite driver first:
//   go get github.com/mattn/go-sqlite3
