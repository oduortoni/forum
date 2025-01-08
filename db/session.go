package db

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize the database (create the table) if not already set up.
func InitializeSessionTable() {
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			token TEXT PRIMARY KEY,
			username TEXT NOT NULL
		)
	`)
	if err != nil {
		fmt.Println("Error creating table:", err)
	}
}

// Remember retrieves the username associated with the provided token.
func Remember(token string) string {
	var username string
	err = db.QueryRow("SELECT username FROM sessions WHERE token = ?", token).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		fmt.Println("Error executing query:", err)
		return ""
	}

	return username
}

// Save generates a new token and saves the username-token pair to the database.
func Save(username string, token string) (string, bool) {
	_, err = db.Exec("INSERT INTO sessions (token, username) VALUES (?, ?)", token, username)
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return "", false
	}

	return token, true
}

// Delete removes the session associated with the provided token.
func Delete(token string) bool {
	result, err := db.Exec("DELETE FROM sessions WHERE token = ?", token)
	if err != nil {
		fmt.Println("Error deleting from database:", err)
		return false
	}

	// check if any rows were deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error checking affected rows:", err)
		return false
	}

	return rowsAffected > 0
}

// generateToken generates a random 16-character token.
func GenerateToken() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 16)

	for i := range token {
		token[i] = characters[rand.Intn(len(characters))]
	}

	return string(token)
}
