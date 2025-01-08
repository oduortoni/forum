package db

import (
	"database/sql"
	"fmt"
	"log"

	"forum/models"
)

func InitializeUsersTable() {
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			username TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterUser(email, username, password string) bool {
	stmt, err := db.Prepare("INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)")
	if err != nil {
		// "Internal server error" => http.StatusInternalServerError
		fmt.Printf("[ERROR](db.RegisterUser)||Internal server error||: %s\n", err.Error())
		return false
	}

	_, err = stmt.Exec(email, username, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			// "Email already taken" => ttp.StatusConflict
			fmt.Printf("[ERROR](db.RegisterUser)||UNIQUE constraint failed||: %s\n", err.Error())
		} else {
			// "Internal server error" => http.StatusInternalServerError
			fmt.Printf("[ERROR](db.RegisterUser)||Internal server error||: %s\n", err.Error())
		}
		return false
	}
	return true
}

func GetUserByEmail(email string) *models.User {
	user := &models.User{}
	err := db.QueryRow("SELECT id, email, username, password_hash FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// "Invalid email or password" => http.StatusUnauthorized
			fmt.Printf("[ERROR](db.GetUserByEmail)||Invalid email or password||: %s\n", err.Error())
		} else {
			fmt.Printf("[ERROR](db.GetUserByEmail)||Internal server error||: %s\n", err.Error())
			// "Internal server error" => http.StatusInternalServerError
		}
		return nil
	}
	return user
}
