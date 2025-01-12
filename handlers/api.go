package handlers

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/pkg/api/apiVersions"
)

func ApiVersionsController(conn net.Conn, request []byte) {
	response, err := apiVersions.ParseRequest(conn, request)
	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		return
	}

	if _, err := conn.Write(response); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}

	fmt.Println("Processed a request and waiting for the next")
}

func DescribeTopicPartitionsController(conn net.Conn, request []byte) {
	response, err := apiVersions.ParseRequest(conn, request)
	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		return
	}

	if _, err := conn.Write(response); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}

	fmt.Println("Processed a request and waiting for the next")
}