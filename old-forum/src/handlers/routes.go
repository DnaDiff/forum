package handlers

import (
	"database/sql"
	"net/http"
)

func RouteHandler(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", createHandlerFunc(HandleIndex, db))
	// mux.HandleFunc("/login", createHandlerFunc(HandleLogin, db))
	// mux.HandleFunc("/logout", createHandlerFunc(HandleLogout, db))
	// mux.HandleFunc("/register", createHandlerFunc(HandleRegister, db))
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/assets/img/coc.png")
	})
	return mux
}

func createHandlerFunc(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
