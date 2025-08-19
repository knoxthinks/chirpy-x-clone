package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()

	// Serve index.html and other root files
	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/", fileServer)

	// Serve files under /assets/ from ./assets/ directory
    mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}