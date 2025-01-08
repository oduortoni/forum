package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize the database (create the table) if not already set up.
func InitializeSessionTable() {
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			token TEXT PRIMARY KEY,
			userid INTEGER NOT NULL,
			username TEXT NOT NULL
		)
	`)
	if err != nil {
		fmt.Println("Error creating table:", err)
	}
}

// Remember retrieves the username associated with the provided token.
func Remember(token string) (int, string) {
	var userID int = 0
	var username string
	err = db.QueryRow("SELECT userid, username FROM sessions WHERE token = ?", token).Scan(&userID, &username)
	if err != nil {
		if err == sql.ErrNoRows {
			return userID, ""
		}
		fmt.Println("Error executing query:", err)
		return userID, ""
	}

	return userID, username
}

// Save generates a new token and saves the username-token pair to the database.
func Save(userid int, username string, token string) (string, bool) {
	_, err = db.Exec("INSERT INTO sessions (token, userid, username) VALUES (?, ?, ?)", token, userid, username)
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
