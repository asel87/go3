package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/User/go/pkg/store/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5433"
	user     = "postgres"
	password = "asel3127"
	dbname   = "postgres"
)

func main() {
	// Build the PostgreSQL connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	// Your code using the connected database...

	// Close the database connection when done
	defer db.Close()
}
