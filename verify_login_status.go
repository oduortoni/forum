package main

import (
	"fmt"
	"net/http"

	"forum/db"
)

// Middleware to check if user is logged in
func VerifyLoginStatus(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie(SESSIONCOOKIENAME)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}

	// Validate session (cookie)
	sessionID := cookie.Value
	var userID int
	err = db.DB.QueryRow("SELECT id FROM users WHERE id = ?", sessionID).Scan(&userID)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}

	return true
}
