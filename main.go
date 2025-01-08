package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/db"
	"forum/controllers"
	"forum/controllers/posts"
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

	http.HandleFunc("/posts/create", posts.Post)

	log.Fatal(http.ListenAndServe(address, nil))
}

// package main

// import (
// 	"fmt"
// 	"forum/db"
// )

// func init() {
// 	db.Init() // initialize the database connection
// }

// func main() {
// 	defer db.Close()

// 	ids := []int{1, 3, 5}
// 	id, err := db.CreatePost(1, "Extra Pressure", "A kenyan song that is kool by nature", ids)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println("ID: ", id)
// }
