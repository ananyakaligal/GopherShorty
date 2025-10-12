package main

import (
	"fmt"
	"log"
	"net/http"
)

// handlerRoot is the handler for the root path "/"
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	// Fprintf writes a formatted string to a writer.
	// Here, w (the ResponseWriter) is our writer, so the text goes to the client.
	fmt.Fprintf(w, "Welcome to GopherShorty!")
}

func main() {
	// http.HandleFunc registers a function to handle a specific path.
	// Any request to "/" will be handled by our handlerRoot function.
	http.HandleFunc("/", handlerRoot)

	// Define the port we want to listen on.
	port := ":8080"
	log.Printf("Starting server on port %s\n", port)

	// http.ListenAndServe starts the server.
	// It blocks forever, unless it encounters a fatal error.
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
