package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DnaDiff/forum/old-forum/src/errors"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" {
		log.Printf("Wrong path: %s\n", r.URL.Path)
		errors.HttpError(w, errors.Error{Code: http.StatusNotFound})
		return
	}

	switch r.Method {
	case http.MethodGet:
		ExecuteTemplate(w, "index.html", nil)
	default:
		log.Println("Wrong method")
		errors.HttpError(w, errors.Error{Code: http.StatusMethodNotAllowed})
		return
	}

}
