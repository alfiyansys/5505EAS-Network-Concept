package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	// 1. Establish a listener for incoming TCP connections.
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()

	fmt.Printf("TCP Server listening on %s:%s\n", HOST, PORT)
	fmt.Println("--------------------------------------------")

	// 2. Main loop to accept new connections.
	for {
		// listener.Accept() blocks until a new client connects.
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// 3. Handle the connection concurrently using a goroutine.
		// This is the key to handling multiple, simultaneous requests.
		go handleConnection(conn)
	}
}

// handleConnection processes the individual client connection.
func handleConnection(conn net.Conn) {
	// Ensure the connection is closed when the goroutine finishes.
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	log.Printf("New connection from %s", clientAddr)

	// Set a deadline for reading to prevent connection abuse/hanging.
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	// Use a buffered reader to read the entire request line by line.
	scanner := bufio.NewScanner(conn)

	// Read the first line (the Request Line: GET / HTTP/1.1)
	if scanner.Scan() {
		requestLine := scanner.Text()

		// Log the request line for demonstration
		log.Printf("[%s] Request: %s", clientAddr, requestLine)

		// --- Simple Request Parsing (Optional, but useful for a "web server") ---
		// Split the request line to get the method and path
		parts := strings.Fields(requestLine)
		var path string
		if len(parts) >= 2 {
			path = parts[1] // The path is the second element (e.g., /index.html)
		}

		// 4. Construct the HTTP Response manually.
		// Even though we are using raw sockets, to be a "web server,"
		// we must speak the HTTP protocol.

		// Status Line
		statusLine := "HTTP/1.1 200 OK\r\n"

		// Headers
		contentType := "Content-Type: text/plain; charset=utf-8\r\n"
		dateHeader := fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123))

		// Body Content
		bodyContent := fmt.Sprintf("Hello from the raw TCP Go Server!\nRequested Path: %s\n", path)
		contentLength := fmt.Sprintf("Content-Length: %d\r\n", len(bodyContent))

		// Full Response Message
		response := statusLine + contentType + contentLength + dateHeader + "\r\n" + bodyContent

		// 5. Send the response back to the client.
		_, err := conn.Write([]byte(response))
		if err != nil {
			log.Printf("[%s] Error writing response: %v", clientAddr, err)
		}
	}

	if err := scanner.Err(); err != nil {
		// This happens if the client disconnects or the deadline is reached.
		log.Printf("[%s] Error reading from connection: %v", clientAddr, err)
	}

	log.Printf("Connection closed for %s", clientAddr)
}
