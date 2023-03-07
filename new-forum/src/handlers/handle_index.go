package handlers

import (
	"database/sql"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	http.ServeFile(w, r, "./public/index.html")
}
