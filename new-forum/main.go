package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
	"github.com/DnaDiff/forum/new-forum/src/handlers"
	_ "github.com/mattn/go-sqlite3"
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
