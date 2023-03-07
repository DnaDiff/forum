package errors

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func HttpError(w http.ResponseWriter, err Error) {
	// If the error code is not set, set it to 500
	if err.Code == 0 {
		err.Code = http.StatusInternalServerError
	}
	// If the error message is not set, set it to the default message for the error code
	if err.Message == "" {
		err.Message = http.StatusText(err.Code)
	}
	// If the original error is not set, set it to the error message
	if err.Original == nil {
		err.Original = errors.New(err.Message)
	}

	// Print the error to the console and send the error to the client
	fmt.Fprintln(os.Stderr, err.Code, err.Original)
	// http.Error(w, err.Message, err.Code)
	handleError(w, &err)
}

func handleError(w http.ResponseWriter, err *Error) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(err.Code)

	if err := tmpl.ExecuteTemplate(w, "error.html", err); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
