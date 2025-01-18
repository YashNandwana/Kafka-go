package describeTopicPartitions

import (
	"fmt"
)

type ResponseBodyTopicsArray struct {
	ErrorCode					int16
	Topic						TopicMetadata
	TopicID 					[]byte
	IsInternal					int8
	PartitionsArray				PartitionsArray
	TopicAuthorizedOperations	uint32
	TagBuffer					int8
}

func (topicsArray *ResponseBodyTopicsArray) PushToByteSliceResponse() []byte {
	var response []byte
	response = append(response, []byte{0, byte(topicsArray.ErrorCode)}...)
	response = append(response, byte(topicsArray.Topic.TopicNameLength + 1))
	response = append(response, []byte(topicsArray.Topic.TopicName)...)
	response = append(response, topicsArray.TopicID...)
	response = append(response, byte(topicsArray.IsInternal))
	response = append(response, byte(topicsArray.PartitionsArray.PartitionsArrayLength))
	response = append(response, topicsArray.PartitionsArray.PartitionsArrayContent...)
	response = append(response, []byte{0, 0, 0, 0}...)
	response = append(response, []byte{0}...)
	response = append(response, []byte{0xff}...)
	response = append(response, []byte{0}...)
	return response
}

func GenerateTopicsArrayResponse(topicsMetadataSlice []TopicMetadata) []byte {
	var (
		response []byte
		topicsArray ResponseBodyTopicsArray
	)
	response = append(response, []byte{byte(len(topicsMetadataSlice) + 1)}...)
	fmt.Println(len(topicsMetadataSlice) + 1)
	
	for _, topicMetadata := range topicsMetadataSlice {
		topicsArray.ErrorCode = 3
		topicsArray.Topic = topicMetadata
		topicsArray.TopicID = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		topicsArray.IsInternal = 0
		var pa PartitionsArray
		pa.PartitionsArrayLength = 1
		topicsArray.PartitionsArray	= pa
	}
	
	bytifiedTopicsArray := topicsArray.PushToByteSliceResponse()
	response = append(response, bytifiedTopicsArray...)

	return response
}