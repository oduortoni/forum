package db

import (
	"fmt"
	"log"
	"strconv"

	"forum/models"
)


var Categories = []models.Category{
	{ID: 0, Name: "All"},
	{ID: 1, Name: "Education"},
	{ID: 2, Name: "Science"},
	{ID: 3, Name: "Technology"},
	{ID: 4, Name: "Culture"},
	{ID: 5, Name: "Sports"},
	{ID: 6, Name: "History"},
	{ID: 7, Name: "Miscellenous"},
}

// Initialize the database with posts and categories tables
func InitializePostTables() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else { // populate the table for categories
		for _, category := range Categories {
			CreateCategory(category.ID, category.Name)
		}
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS post_categories (
			post_id INTEGER NOT NULL,
			category_id INTEGER NOT NULL,
			FOREIGN KEY(post_id) REFERENCES posts(id),
			FOREIGN KEY(category_id) REFERENCES categories(id),
			PRIMARY KEY(post_id, category_id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// Create a new post by a user
func CreatePost(userID int, title, content string, categoryIDs []int) (int, error) {
	// Insert the post into the database
	stmt, err := db.Prepare("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	res, err := stmt.Exec(userID, title, content)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %v", err)
	}

	// Get the ID of the newly created post
	postID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	// Associate categories with the post
	for _, categoryID := range categoryIDs {
		_, err = db.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postID, categoryID)
		if err != nil {
			return 0, fmt.Errorf("failed to associate post with category: %v", err)
		}
	}

	return int(postID), nil
}

// Get all posts with their associated categories
func GetAllPosts() ([]models.Post, error) {
	rows, err := db.Query(`
		SELECT p.id, p.title, p.content, p.created_at, u.username 
		FROM posts p
		JOIN users u ON p.user_id = u.id
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %v", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %v", err)
		}

		// Get categories for each post
		categories, _ := GetCategoriesForPost(post.ID)
		post.Categories = append(post.Categories, categories...) 

		posts = append(posts, post)
	}

	return posts, nil
}

// Get categories associated with a post
func GetCategoriesForPost(postID int) ([]models.Category, error) {
	rows, err := db.Query(`
		SELECT c.id, c.name 
		FROM categories c
		JOIN post_categories pc ON c.id = pc.category_id
		WHERE pc.post_id = ?`, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %v", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var categoryId string
		var categoryName string
		err := rows.Scan(&categoryId, &categoryName)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		cid, err := strconv.Atoi(categoryId)
		if err != nil {
			return nil, fmt.Errorf("failed to convert categoryId to integer: %v", err)
		}

		category := models.Category{
			ID:   cid,
			Name: categoryName,
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetAllCatogories() []models.Category {
	return Categories
}

// Create a category
func CreateCategory(id int, name string) error {
	stmt, err := db.Prepare("INSERT INTO categories (id, name) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	_, err = stmt.Exec(id, name)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}
