package describeTopicPartitions

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/utils"
	"github.com/davecgh/go-spew/spew"
)

func ParseRequest(request []byte) ([]byte, error) {
	fmt.Println(request)
	response := make([]byte, 4)

	// parse correlation_id
	correlationId := request[8:12]
	fmt.Printf("correlationId: ")
	fmt.Println(utils.ConvertToInt16(correlationId))

	// parse clientId - length and contents
	clientIdLength := request[12:14]
	fmt.Printf("clientLength: ")
	fmt.Println(utils.ConvertToInt16(clientIdLength))

	clientIdContent := request[14:14 + utils.ConvertToInt16(clientIdLength)]
	fmt.Printf("clientIdContent: ")
	fmt.Println(string(clientIdContent))

	// tag buffer index
	tagBufferIndex := 14 + utils.ConvertToInt16(clientIdLength)
	requestBodyStartIndex := tagBufferIndex + 1

	topicsArrayLength := request[requestBodyStartIndex:requestBodyStartIndex + 1]
	fmt.Printf("topicsArrayLength: ")
	fmt.Println(utils.ConvertToInt8(topicsArrayLength))

	topicsCount := utils.ConvertToInt8(topicsArrayLength) - 1
	topicsMetadataSlice, currentIndex := GetAllTopics(request, topicsCount, requestBodyStartIndex + 1)
	spew.Dump(topicsMetadataSlice)

	// response partition limit
	responsePartitionLimit := request[currentIndex: currentIndex + 4]
	fmt.Printf("responsePartitionLimit: ")
	fmt.Println(utils.ConvertToInt8(responsePartitionLimit))

	// add correlation id
	response = append(response, correlationId...)
	response = append(response, byte(utils.ConvertToInt8([]byte{0})))
	
	// add throttle time to response
	throttleTime := []byte{0, 0, 0, 0}
	response = append(response, throttleTime...)

	topicsArrayResponse := GenerateTopicsArrayResponse(topicsMetadataSlice)
	response = append(response, topicsArrayResponse...)

	response = AddMessageSizeToResponse(response)

	return response, nil
	
}
