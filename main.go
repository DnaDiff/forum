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

	// CreatePost(db, &Post{UserID: 1, Title: "Hello, world!", Content: "This is my first post!", Category: "General"})

	// CreateComment(db, &Comment{PostID: 1, UserID: 1, Content: "Hello, world!"})

	// post, err := GetUserPosts(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, p := range post {
	// 	log.Println(p)
	// }

	// nice, err := GetUserComments(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, p := range nice {
	// 	log.Println(p)
	// }

	// nice, err := GetUserLikes(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(nice)

	// RemovePost(db, 1)

}
