package controllers

import (
	"fmt"
	"net/http"

	"forum/db"
)

// Middleware to check if user is logged in
func Authorize(w http.ResponseWriter, r *http.Request) (int, string) {
	cookie, err := r.Cookie(SESSIONCOOKIENAME)
	if err != nil {
		fmt.Println(err)
		// http.Redirect(w, r, "/login", http.StatusFound)
		return 0, ""
	}

	// Validate session (cookie)
	sessionID := cookie.Value
	return db.Remember(sessionID)
}

func Revoke(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie(SESSIONCOOKIENAME)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}

	// Validate session (cookie)
	sessionID := cookie.Value
	return db.Delete(sessionID)
}

