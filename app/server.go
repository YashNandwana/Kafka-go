package main

import (
	"fmt"
	"net"
	"os"
	"github.com/codecrafters-io/kafka-starter-go/handlers"

)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")
	
	listener, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092: ", err.Error())
		os.Exit(1)
	}

	if err := handlers.TcpController(listener); err != nil {
		fmt.Printf("Error in TcpController: %v\n", err)
		os.Exit(1)
	}
	
}
