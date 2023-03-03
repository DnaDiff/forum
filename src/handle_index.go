package forum

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/DnaDiff/forum/src/errors"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" {
		fmt.Println("Wrong path: " + r.URL.Path)
		errors.HttpError(w, errors.Error{Code: http.StatusNotFound})
		return
	}

	if r.Method == "GET" {
		ExecuteTemplate(w, "index.html", nil)
	} else {
		fmt.Println("Wrong method")
		errors.HttpError(w, errors.Error{Code: http.StatusMethodNotAllowed})
		return
	}
}
