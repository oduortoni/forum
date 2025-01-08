package posts

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"forum/controllers"
	"forum/db"
	"forum/models"
)

// Dashboard handler that shows the form for creating posts and displays all posts
func Post(w http.ResponseWriter, r *http.Request) {
	userID, username := controllers.Authorize(w, r)
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Get available categories
	categories := db.GetAllCatogories()

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		categoryIDs := []int{} // Parse selected category IDs from form

		// Parse the selected category IDs from the form
		for _, id := range r.Form["categories"] {
			// Convert string ID to integer
			categoryID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Fprintf(w, "Error parsing category ID: %s", err)
				return
			}
			categoryIDs = append(categoryIDs, categoryID)
		}

		// Create post
		_, err := db.CreatePost(userID, title, content, categoryIDs)
		if err != nil {
			fmt.Fprintf(w, "Error creating post: %s", err)
			return
		}

		// fmt.Fprintf(w, "Post created successfully with ID %d", postID)
		http.Redirect(w, r, "/posts/create", http.StatusFound)
		return
	}

	// Display all posts
	posts, err := db.GetAllPosts()
	if err != nil {
		fmt.Fprintf(w, "Error fetching posts: %s", err)
		return
	}

	// Render the dashboard template
	tmpl := template.Must(template.ParseFiles("templates/posts/create.html"))
	tmpl.Execute(w, struct {
		Username   string
		Categories []models.Category
		Posts      []models.Post
	}{
		Username:   username,
		Categories: categories,
		Posts:      posts,
	})
}
