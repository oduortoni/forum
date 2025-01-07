package main

import (
	"fmt"
	"forum/db"
	"log"
	"net/http"
)

func init() {
	db.Init() // initialize the database connection
}

func main() {
	defer db.Close() // close the db conn after application terminates

	port := 9000

	// Set up routes
	http.HandleFunc("/", Index)
	http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/register", RegisterView)
	http.HandleFunc("/register/submit", RegisterController)
	http.HandleFunc("/login", LoginView)
	http.HandleFunc("/login/submit", LoginController)
	http.HandleFunc("/logout", Logout)

	fmt.Printf("Server running on http://localhost:%d\n", port)
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
