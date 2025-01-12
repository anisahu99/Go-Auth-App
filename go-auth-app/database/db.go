package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var DB *sql.DB

// Connect initializes the connection to the MySQL database
func Connect() {
	var err error

	// Load the connection string from environment variables
	dsn := os.Getenv("DB_CONNECTION_STRING")
	if dsn == "" {
		log.Fatal("DB_CONNECTION_STRING is not set in the environment variables")
	}

	// Open a connection to the MySQL database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	// Create the users table if it doesn't exist
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL
        );
    `)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Connected to the MySQL database successfully!")
}
