package controllers

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"forum/db"
	"forum/utils"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user := db.GetUserByEmail(email)
		if user == nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Compare password
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Create session (cookie)
		sessionID := utils.GenerateToken(0)
		db.Save(user.ID, user.Username, sessionID)
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
