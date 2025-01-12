package handlers

import (
	"encoding/binary"
	"fmt"
	"net"

)

type ApiHandler func(conn net.Conn, request []byte)

var apiHandlersMap = map[int16]ApiHandler {
	18: ApiVersionsController,
	75: DescribeTopicPartitionsController,
}

func TcpController(listener net.Listener) error {
	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue // Log the error and continue to the next iteration
		}

		fmt.Println("Accepted a connection. Listening for requests on 0.0.0.0:9092")

		// Handle the connection in a separate goroutine for concurrency
		go handleConnectionRequests(conn)
	}
}

// handleConnectionRequests processes requests from a single client over a persistent connection.
func handleConnectionRequests(conn net.Conn) {
	
	fmt.Println("Handling a new connection")

	buff := make([]byte, 1024)

	// Loop to handle multiple sequential requests
	for {
		// Read data from the client
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("Connection closed or error reading data: %v\n", err)
			return
		}

		// remove the invalid part from buff
		request := buff[:n]
		apiKey := fetchApiKeyFromResponse(request)
		apiHandlersMap[apiKey](conn, request)
	}
}

func fetchApiKeyFromResponse(request []byte) int16 {
	return int16(binary.BigEndian.Uint16(request[4:6]))
}