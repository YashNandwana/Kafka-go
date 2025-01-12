package describeTopicPartitions

import (
	"fmt"
	"net"
)

func ParseRequest(conn net.Conn, request []byte) ([]byte, error) {
	resp := make([]byte, 4)

	// parse correlation_id
	correlationId := request[4:8]
	fmt.Println("correlationId: " + string(correlationId))
	return resp, nil
	
}