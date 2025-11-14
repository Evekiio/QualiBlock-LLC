package logic

import "net/http"

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	// Get the Incoming Request Path
	path := r.URL.Path

	// Respond with a simple message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to QualiBlock Web Platform! You requested: " + path))
}
