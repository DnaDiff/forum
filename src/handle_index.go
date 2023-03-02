package forum

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/DnaDiff/forum/src/errors"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == "GET" {
		ExecuteIndex(w, "index.html", nil)
	}
}

func ExecuteIndex(w http.ResponseWriter, file string, data interface{}) {
	templates, err := filepath.Glob("templates/*.html")
	if err != nil {
		errors.HttpError(w, errors.Error{Code: http.StatusInternalServerError, Message: errors.TEMPLATE_CORRUPTED_ERROR, Original: err})
		return
	}

	tmpl := template.Must(template.New("index.html").ParseFiles(templates...))

	err = tmpl.ExecuteTemplate(w, file, data)
	if err != nil {
		errors.HttpError(w, errors.Error{Code: http.StatusInternalServerError, Message: errors.TEMPLATE_CORRUPTED_ERROR, Original: err})
		return
	}
}
