package main

import (
	"net/http"

	"forum/db"

	"golang.org/x/crypto/bcrypt"
)

// Handler to display the registration page
func RegisterView(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/register.html")
}

// Handler for user registration
func RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Hash the password
		Password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Insert the user into the database
		stmt, err := db.DB.Prepare("INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec(email, username, Password)
		if err != nil {
			if err.Error() == "UNIQUE constraint failed: users.email" {
				http.Error(w, "Email already taken", http.StatusConflict)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
