package main

import (
	"database/sql"
)

var database *sql.DB

func init() {
	// Create DB
	database, _ = sql.Open("sqlite3", "../data/db/user_data.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, UUID TEXT, emails TEXT, date TEXT, pincode TEXT, ageGroup INTEGER)")
	statement.Exec()

	// Create workers for users already in DB
	// initDB()
}
