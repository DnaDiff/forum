package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/DnaDiff/forum/new-forum/src/handlers"
)

const PORT = "8080"

func main() {
	// Establish connection to the database
	// var db = database.ConnectDB()
	var db *sql.DB

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
