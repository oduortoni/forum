package models

type Post struct {
	ID         int
	Title      string
	Content    string
	CreatedAt  string
	Username   string
	Categories []Category
}

type Category struct {
	ID   int
	Name string
}
