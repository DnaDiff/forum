package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

// var store = sessions.NewCookieStore([]byte("my_secret_key"))

func HandleLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
}
