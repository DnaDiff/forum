package handlers

import (
	"database/sql"
	"net/http"
)

var mux *http.ServeMux

func RouteHandler(db *sql.DB) *http.ServeMux {
	mux = http.NewServeMux()
	mux.HandleFunc("/", createHandlerFunc(HandleIndex, db))                  // Main page
	mux.HandleFunc("/api/categories", createHandlerFunc(HandleContent, db))  // Returns all categories
	mux.HandleFunc("/api/categories/", createHandlerFunc(HandleContent, db)) // Category and post specific requests, e.g. data, upvote, downvote, comment, etc.
	mux.HandleFunc("/api/login", createHandlerFunc(HandleLogin, db))
	// mux.HandleFunc("/logout", createHandlerFunc(HandleLogout, db))
	mux.HandleFunc("/api/register", createHandlerFunc(HandleRegister, db))
	return mux
}

func createHandlerFunc(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
