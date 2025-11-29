package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// The server's address and port (adjust if necessary based on your server setup)
const (
	serverIP   = "127.0.0.1" // Localhost
	serverPort = "12000"     // Port used in the Kurose & Ross example
	numPings   = 10
	timeoutSec = 1
)

func main() {
	// 1. Resolve the server address
	serverAddr, err := net.ResolveUDPAddr("udp", serverIP+":"+serverPort)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	// 2. Dial the server
	// In Go's net package, DialUDP creates a connection-oriented UDP socket,
	// which is ideal for a client talking to a single server.
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error dialing server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("UDP Pinger Client started. Pinging %s:%s...\n", serverIP, serverPort)

	// Set the read timeout for the connection
	// This implements the requirement to wait up to one second for a reply.
	err = conn.SetReadDeadline(time.Now().Add(timeoutSec * time.Second))
	if err != nil {
		fmt.Println("Error setting read deadline:", err)
		os.Exit(1)
	}

	// 3. Send 10 ping messages
	for i := 1; i <= numPings; i++ {
		// Prepare the ping message
		// Format: "Ping <sequence_number> <timestamp>"
		// The server is expected to echo this back.
		now := time.Now().UnixMilli() // Current time in milliseconds
		message := fmt.Sprintf("Ping %d %d", i, now)

		// Record the time just before sending
		startTime := time.Now()

		// Send the ping message
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending ping:", err)
			continue
		}

		// Buffer to hold the incoming pong message
		buffer := make([]byte, 1024)

		// 4. Wait for the reply (up to 1 second due to SetReadDeadline)
		n, _, err := conn.ReadFromUDP(buffer)

		if err != nil {
			// Check if the error is a timeout
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// Packet was lost/timed out
				fmt.Printf("PING %d: Request timed out\n", i)
			} else {
				// Some other read error
				fmt.Printf("PING %d: Read error: %v\n", i, err)
			}
		} else {
			// Pong received successfully
			endTime := time.Now()
			// Calculate RTT in milliseconds
			rtt := endTime.Sub(startTime).Milliseconds()

			// The server's reply is in the buffer
			response := string(buffer[:n])

			// Print the RTT
			fmt.Printf("PING %d: Received '%s', RTT: %d ms\n", i, response, rtt)
		}

		// Wait a moment before sending the next ping (optional, but good practice)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\nUDP Pinger Client finished.")
}
