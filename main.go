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

	// LikePost(db, 1, 1)
	// DislikePost(db, 1, 2)
	// RemoveDislikePost(db, 1, 2)
	// RemoveLikePost(db, 1, 2)

	// CreatePost(db, &Post{ // Create a post
	// 	UserID:   1,
	// 	Title:    "My fourth post",
	// 	Content:  "This is my fourth post",
	// 	Category: "General",
	// })

	// RemovePost(db, 1) // Remove a post

	// u := User{
	// 	ProfilePicture: "mssdasa",
	// 	Username:       "johnsasddsfasd",
	// 	Age:            25,
	// 	Gender:         "male",
	// 	FirstName:      "Jo",
	// 	LastName:       "Smi",
	// 	Password:       "password456",
	// 	Email:          "JohnsdfsSmdashe@example.com",
	// }
	// if !CheckDuplicateUsername(db, u.Username) {

	// 	if err := CreateUser(db, &u); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// u, err := database.GetUserByUsername(db, "johnsads")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("User: %+v \n", u)

	// database.DeleteUserByUsername(db, "johns")

}
