package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/db"
	"forum/controllers"
)

func init() {
	db.Init() // initialize the database connection
}

func main() {
	defer db.Close() // close the db conn after application terminates

	port := 9000

	// Set up routes
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/dashboard", controllers.Dashboard)
	http.HandleFunc("/register", controllers.RegisterView)
	http.HandleFunc("/register/submit", controllers.RegisterController)
	http.HandleFunc("/login", controllers.LoginView)
	http.HandleFunc("/login/submit", controllers.LoginController)
	http.HandleFunc("/logout", controllers.Logout)

	fmt.Printf("Server running on http://localhost:%d\n", port)
	address := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
