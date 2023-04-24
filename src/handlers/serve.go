package handlers

import (
	"net/http"
	"path"
)

func ServeFiles(mux *http.ServeMux) {
	serveDir(mux, "./public/assets/fonts")
	serveDir(mux, "./public/css")
	serveDir(mux, "./public/js")
	serveDir(mux, "./public/assets/images")
}

func serveDir(mux *http.ServeMux, dirPath string) {
	dir := "/" + path.Base(dirPath) + "/"
	fs := http.FileServer(http.Dir(dirPath))
	mux.Handle(dir, (http.StripPrefix(dir, fs)))
}
