package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./front-end/templates/home.html")
	})
	// Serve CSS files
	cssHandler := http.StripPrefix("/css/", http.FileServer(http.Dir("./front-end/css/")))
	http.Handle("/css/", cssHandler)

	// Serve JS files
	jsHandler := http.StripPrefix("/js/", http.FileServer(http.Dir("./back-end/JS/")))
	http.Handle("/js/", jsHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
