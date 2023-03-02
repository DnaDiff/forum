package errors

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func HttpError(w http.ResponseWriter, err Error) {
	if err.Code > 0 && err.Message == "" {
		err.Message = http.StatusText(err.Code)
	} else if err.Code == 0 && err.Message == "" {
		err.Code = http.StatusInternalServerError
		err.Message = http.StatusText(err.Code)
	}
	fmt.Fprintln(os.Stderr, err.Original)
	handleError(w, &err)
}

func handleError(w http.ResponseWriter, err *Error) {
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	w.WriteHeader(err.Code)
	tmpl.ExecuteTemplate(w, "error.html", err)
}
