package handlers

import (
	"database/sql"
	"net/http"
)

var mux *http.ServeMux

func RouteHandler(db *sql.DB) *http.ServeMux {
	mux = http.NewServeMux()
	mux.HandleFunc("/", createHandlerFunc(HandleIndex, db))
	mux.HandleFunc("/api/posts", createHandlerFunc(RetrievePosts, db))
	// mux.HandleFunc("/login", createHandlerFunc(HandleLogin, db))
	// mux.HandleFunc("/logout", createHandlerFunc(HandleLogout, db))
	// mux.HandleFunc("/register", createHandlerFunc(HandleRegister, db))
	return mux
}

func createHandlerFunc(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
