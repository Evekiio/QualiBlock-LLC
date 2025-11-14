package logic

import (
	"net/http"
	"os"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	// Get incoming request path
	path := r.URL.Path

	// Root route serves index.html
	if path == "/" {
		http.ServeFile(w, r, "./presentation/public/index.html")
		return
	}

	// Trim leading slash and build expected HTML file path
	fileName := path[1:] + ".html"
	filePath := "./presentation/public/" + fileName

	// Attempt to serve the file
	if _, err := os.Stat(filePath); err == nil {
		http.ServeFile(w, r, filePath)
		return
	}

	// If not found, serve 404 error page
	http.ServeFile(w, r, "./presentation/public/404.html")
}
