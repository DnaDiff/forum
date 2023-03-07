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
	// database.AddUser(db)

	// Get a user from the database
	// user, err := database.GetUser(db)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// userInfo, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(userInfo) + "\n" + "User read successfully.")

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
