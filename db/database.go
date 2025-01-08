package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

const (
	dbFile = "./data/forum.db"
)

// Initialize the database connection
func Init() {
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	InitializeUsersTable()
	InitializeSessionTable()
	InitializePostTables()
}

// Close the database connection
func Close() {
	db.Close()
}
