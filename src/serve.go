package forum

import (
	"net/http"
	"path"
)

func ServeFiles() {
	serveDir("../public/assets/fonts")
	serveDir("../public/css")
	serveDir("../public/js")
	serveDir("../public/assets/img")
	serveDir("../templates")
}

func serveDir(dirPath string) {
	dir := "/" + path.Base(dirPath) + "/"
	fs := http.FileServer(http.Dir(dirPath))
	http.Handle(dir, (http.StripPrefix(dir, fs)))
}
