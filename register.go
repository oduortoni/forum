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
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if ok := db.RegisterUser(email, username, string(hashedPassword)); !ok {
			http.Error(w, "Unable to register user", http.StatusConflict)
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
