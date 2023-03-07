package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/DnaDiff/forum/old-forum/src/errors"
)

// Execute a page or template to render on index.html
func ExecuteTemplate(w http.ResponseWriter, file string, data interface{}) {
	log.Printf("Parsing and executing the template %s", file)

	tmpl := template.Must(template.New("index.html").ParseGlob("templates/*.html"))

	err := tmpl.ExecuteTemplate(w, file, data)
	if err != nil {
		errors.HttpError(w, errors.Error{Code: http.StatusInternalServerError, Message: errors.TEMPLATE_CORRUPTED_ERROR, Original: err})
		return
	}
}
