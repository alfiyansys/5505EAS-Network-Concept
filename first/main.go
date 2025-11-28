package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 1. Define the routes (URL paths) and associate them with handler functions.

	// The default route ("/") uses a simple inline function (closure) as the handler.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w (http.ResponseWriter) is what you use to write the response back to the client.
		// r (*http.Request) contains all the details of the incoming request.
		fmt.Fprintf(w, "Welcome to the Go Web Server Skeleton!")
	})

	// A more specific route ("/hello") uses a separate, named handler function.
	http.HandleFunc("/hello", helloHandler)

	// 2. Start the server and specify the port to listen on.
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)

	// ListenAndServe starts an HTTP server with a given address and handler.
	// Passing 'nil' for the handler uses the default 'http.ServeMux' (the router
	// where we registered our HandleFuncs).
	// log.Fatal is used to ensure the program exits if the server cannot start.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

// Handler function for the "/hello" route.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Simple check to ensure we only respond to GET requests (optional but good practice)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Write the response header and body
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from the Go server!"))
}