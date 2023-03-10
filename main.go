package main

import (
	"database/sql"
	"fmt"
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

	// u := &database.User{
	// 	ProfilePicture: "https://example.com/profile.jpg",
	// 	Username:       "johns",
	// 	Age:            25,
	// 	Gender:         "male",
	// 	FirstName:      "Jo",
	// 	LastName:       "Smi",
	// 	Password:       "password456",
	// 	Email:          "JohnSmhe@example.com",
	// }

	// if err := database.CreateUser(db, u); err != nil {
	// 	log.Fatal(err)
	// }

	u, err := database.GetUserByUsername(db, "johns")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User: %+v \n", u)

}
