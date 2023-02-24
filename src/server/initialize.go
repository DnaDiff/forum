package server

import (
	"fmt"
	"net/http"
	"path"
)

const PORT = "8080"

func StartServer() {
	// Serve static files in optimal order
	serveDir("./public/assets/fonts")
	serveDir("./public/css")
	serveDir("./public/js")
	serveDir("./public/assets/img")
	serveDir("./templates")
	http.Handle("/", indexHandler())

	fmt.Println("Listening on port :" + PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		panic(err)
	}

}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
}

func serveDir(dirPath string) {
	dir := path.Base(dirPath)
	stripped := http.StripPrefix("/"+dir+"/", http.FileServer(http.Dir(dirPath)))
	http.Handle("/"+dir+"/", stripped)
}
