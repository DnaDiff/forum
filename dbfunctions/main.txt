package main

import (
	"database/sql"
	"log"

	// "net/http"
	// "time"

	. "github.com/DnaDiff/forum/src/database"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/DnaDiff/forum/src/handlers"
)

const PORT = "8080"

func main() {
	// Establish connection to the database
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	InitDatabase(db)

	nice, err := GetAllPostsByCategory(db, "General")
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range nice {
		log.Println(post)
	}

}
