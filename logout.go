package main

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	Revoke(w, r);
	
	http.SetCookie(w, &http.Cookie{
		Name:     SESSIONCOOKIENAME,
		Value:    "",
		Expires:  time.Unix(0, 0), // expiry to a time to be in the past
		HttpOnly: true,
		Path:     "/", // must match the path the cookie was set with
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}
