package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	username := Authorize(w, r)
	if username == "" {
		fmt.Fprintf(w, "<h1>User not logged in</h1>")
		// http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, username)
}
