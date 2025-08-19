package main

import (
	"log"
	"net/http"
)
// make a branch for this code

// func main() {
// 	const port = "8080"
// 	const filepathRoot = "."

// 	mux := http.NewServeMux()
	
// 	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//     w.WriteHeader(http.StatusOK)
//     w.Write([]byte("OK"))
// 	})

// 	// Serve index.html and other root files
// 	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

// 	// Serve files under /assets/ from ./assets/ directory
//     mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


// 	srv := &http.Server{
// 		Addr:    ":" + port,
// 		Handler: mux,
// 	}

// 	log.Printf("Serving on port: %s\n", port)
// 	log.Fatal(srv.ListenAndServe())
// }

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", handlerReadiness)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}