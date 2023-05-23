package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	database "github.com/DnaDiff/forum/new-forum/src/dbfunctions"
)

// CreateCookie takes a username and password and creates a cookie for the user
func CreateCookie(w http.ResponseWriter, r *http.Request, db *sql.DB, username string, password string) {

	userID, ok := database.CheckUser(db, username, password)

	if !ok {
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	cookie := http.Cookie{
		Name:   "user",
		Value:  fmt.Sprint(userID),
		Path:   "/",
		MaxAge: 86400, // This cookie will expire after 24 hours
	}

	http.SetCookie(w, &cookie)

	fmt.Fprint(w, "Cookie for user "+username+" created")
}
