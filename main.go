package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/db"
	"forum/controllers"
	"forum/utils"
)

func init() {
	db.Init() // initialize the database connection
}

func main() {
	defer db.Close() // close the db conn after application terminates

	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	address := fmt.Sprintf("0.0.0.0:%d", port)

	// Set up routes
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/dashboard", controllers.Dashboard)
	http.HandleFunc("/register", controllers.RegisterView)
	http.HandleFunc("/register/submit", controllers.RegisterController)
	http.HandleFunc("/login", controllers.LoginView)
	http.HandleFunc("/login/submit", controllers.LoginController)
	http.HandleFunc("/logout", controllers.Logout)

	log.Fatal(http.ListenAndServe(address, nil))
}
