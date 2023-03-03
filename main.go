package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	. "github.com/DnaDiff/forum/src"
)

const PORT = "8080"

var db *sql.DB

func main() {
	// Connect to database

	// Create the multiplexer to handle the routes
	mux := RouteHandler(db)

	// Create the server
	server := http.Server{
		Addr:         ":" + PORT,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	ServeFiles(mux)
	fmt.Println("Listening on port :" + PORT)
	fmt.Println(server.ListenAndServe())
}
