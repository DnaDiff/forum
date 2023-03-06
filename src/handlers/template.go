package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/DnaDiff/forum/src/errors"
)

// Execute a page or template to render on index.html
func ExecuteTemplate(w http.ResponseWriter, file string, data interface{}) {
	log.Printf("Parsing and executing the template %s", file)
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
