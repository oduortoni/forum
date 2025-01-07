package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if !VerifyLoginStatus(w, r) {
		fmt.Fprintf(w, "<h1>User not logged in</h1>")
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, nil)
}
