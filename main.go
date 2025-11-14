package main

import (
	"fmt"
	"net/http"
	"qb/logic"
)

func init() {
	// PLACEHOLDER INIT COMMENT
}

func main() {
	startPlatform()
}

func startPlatform() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register the Logic Routes
	registerRoutes(mux)

	// Display Console Banner
	printBanner()

	fmt.Println("Starting QualiBlock Web Platform on http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error Starting Platform:", err)
	}
}

func printBanner() {
	fmt.Println("")
	fmt.Println("===================================")
	fmt.Println("   QualiBlock LLC - Web Platform  ")
	fmt.Println("===================================")
	fmt.Println("")
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", logic.DefaultHandler)
}
