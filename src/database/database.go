package database

import (
	"database/sql"
	"log"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		panic(err)
	}
	return db
	// create the users table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINRCEMENT, username TEXT, email TEXT, password TEXT, joined DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}

	// create the posts table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, content TEXT, author TEXT, created DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}

}
