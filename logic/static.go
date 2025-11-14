package logic

import "net/http"

func DefaultStaticHandler(w http.ResponseWriter, r *http.Request) {
	// Strip the leading "/static/" prefix and build the local file path
	path := r.URL.Path
	if len(path) < len("/static/") {
		http.NotFound(w, r)
		return
	}

	filePath := "./static/" + path[len("/static/"):]
	http.ServeFile(w, r, filePath)
}
