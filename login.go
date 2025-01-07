package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"forum/db"

	"golang.org/x/crypto/bcrypt"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Find the user by email
		var user User
		err := db.DB.QueryRow("SELECT id, email, username, password_hash FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		// Compare password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Create session (cookie)
		sessionID := fmt.Sprintf("%d", user.ID)
		http.SetCookie(w, &http.Cookie{
			Name:     SESSIONCOOKIENAME,
			Value:    sessionID,
			Expires:  time.Now().Add(24 * time.Hour), // 24 hours
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode, // avoid cross-site cookies
			Path:     "/",
		})

		http.Redirect(w, r, "/dashboard", http.StatusFound)
	}
}
