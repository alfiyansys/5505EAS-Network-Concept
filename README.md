Assignment Repository URL is: https://github.com/alfiyansys/5505EAS-Network-Concept

5505EAS-Network-Concept
===
Written by M Alfiyan Syamsuddin - 1225800008
---

The final semester assignment is to demonstrate how (L7 Application layer) socket programming works and its implementations on programming.

---

1. First Assignment: HTTP Server using RAW TCP Connections

	The first assignment is to demonstrate principles of RAW TCP socket programming and its implementations on HTTP L7 Application protocol layer. It written on Golang instead of python.

	The server is configurated to run on localhost and port 8080. Program begans with listening for incoming connections. Once a connection is established and entered main program loop, the server sends a response to the client using goroutine.

	The client's request is parsed, including IP address, port number, and HTTP method. If the HTTP protocol is present using buffered reader, the server will send a response using buffered writer that manually consutructed HTTP response.

	![RAW TCP based HTTP server access](./assets/http-server.png)

---

2. Second project: UDP Pinger Client Impelementations
	
	The second assignment is to demonstrate how ping messages are implemented using UDP based connections. The server runtime is provided using python. The randominess is to simulate UDP transmission error. 
	
	The provided server code logic shouldn't be modified to preserve assignment constraints. The code is slightly modified to support debugging purposes. 

	Main assignment is to write up and implement pinger client. It written on Golang instead of python at first, until it was changed back to python as instructed.

	
	![Golang based UDP pinger client](./assets/ping-go.png)

	![Python based UDP pinger client](./assets/ping-python.png)

	The Golang based is not working properly after first ping is because of how different programming language implements OS scheduling routine. As server is using python based is using IO blocking model, while Go is non-blocking.

	Based on Pytho3 based UDP ping client, as seen in the image above, the server is running on port 12000. The client is using port 12001. The client is running on localhost, while the server is running on the same machine. Both sides (server and client) are showing in what UNIX time both are communicating, as well as the sequence number.
	
	Test results shows loss rate is 50% due to 0..4 range of loss randominess (loss rate should be 40~50% to be exact, valid). With RTT averaging about 0.3ms, max 0.42ms, min 0.2ms.