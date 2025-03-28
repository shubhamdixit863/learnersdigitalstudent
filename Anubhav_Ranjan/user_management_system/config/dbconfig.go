package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./myusers.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create Users table if not exists
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return DB
}

/***************************************************** */
