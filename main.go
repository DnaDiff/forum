package main

import (
	"database/sql"
	"log"

	// "net/http"
	// "time"

	"github.com/DnaDiff/forum/src/database"
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

	database.InitDatabase(db)

	// make a new user using the CreateUser function
	// make a new post using the CreatePost function
	// make a new comment using the CreateComment function
	// get all posts using the GetAllPosts function
	// get all posts by category using the GetAllPostsByCategory function
	// get all comments by post using the GetAllCommentsByPost function

	// // Create a new user
	// database.CreateUser(db, "mone", "tesdasasd", "dsfsdf@fgd.cesd")

	database.CreateComment(db, 1, 1, "first comment")
	database.CreateComment(db, 1, 1, "second comment")
	database.CreateComment(db, 1, 1, "third comment")

	// // Create a new post
	// database.CreatePost(db, 1, "Nicetest1", "asdasd", "general")
	// database.CreatePost(db, 1, "Nicetest2", "asdasd", "general")
	// database.CreatePost(db, 1, "Nicetest3", "asdasd", "general")

	posts, err := database.GetAllPostsByCategory(db, "general")
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		log.Println(post)
	}

}
