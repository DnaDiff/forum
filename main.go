package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DnaDiff/forum/src/database"
	_ "github.com/mattn/go-sqlite3"

	"github.com/DnaDiff/forum/src/handlers"
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

	

	// database.CreatePost(db, 1, "Title testing 2", "Test content niceasdman")

	// Execute the contents of the init.sql file
	// sqlScript, err := ioutil.ReadFile("./database/init.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = db.Exec(string(sqlScript))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Add a user to the database
	// database.CreateUser(db, "nicebasddasro", "test", " asdsad@cocc.go")

	// Get a user number 1 from the database
	// user, err := database.GetUserByID(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	// Get a user by name
	// user, err := database.GetUserByName(db, "nicebro")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	// Create the mux to handle the routes
	mux := handlers.RouteHandler(db)

	// Create the server
	server := http.Server{
		Addr:         ":" + PORT,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	handlers.ServeFiles(mux)
	fmt.Println("Listening on port :" + PORT)
	fmt.Println(server.ListenAndServe())
}
