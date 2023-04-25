package handlers

import (
	"database/sql"
	"net/http"
)

var mux *http.ServeMux

func RouteHandler(db *sql.DB) *http.ServeMux {
	mux = http.NewServeMux()
	mux.HandleFunc("/", createHandlerFunc(HandleIndex, db))          // Main page
	mux.HandleFunc("/api/posts", createHandlerFunc(HandlePost, db))  // Returns all posts
	mux.HandleFunc("/api/posts/", createHandlerFunc(HandlePost, db)) // Post specific requests, e.g. data, upvote, downvote, comment, etc.
	mux.HandleFunc("/login", createHandlerFunc(HandleLogin, db))
	mux.HandleFunc("/logout", createHandlerFunc(HandleLogout, db))
	mux.HandleFunc("/register", createHandlerFunc(HandleRegister, db))
	mux.HandleFunc("/api/verify-session", createHandlerFunc(HandleVerifySession, db))
	mux.HandleFunc("/api/auth/check", createHandlerFunc(HandleCheckSession, db))


	return mux
}


func createHandlerFunc(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
