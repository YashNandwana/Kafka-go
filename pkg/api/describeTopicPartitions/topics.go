package describeTopicPartitions

import (
	"github.com/codecrafters-io/kafka-starter-go/utils"
)

type TopicMetadata struct {
	TopicNameLength		int16
	TopicName 			string
}

type PartitionsArray struct {
	PartitionsArrayLength	int16
	PartitionsArrayContent	[]byte
}

func GetAllTopics(request []byte, topicsCount int8, topicsArrayFirstIndex int16) ([]TopicMetadata, int16) {
	currentIndex := topicsArrayFirstIndex;
	var topicMetadataSlice []TopicMetadata

	for topicsCount > 0 {
		topicNameLength := int16(utils.ConvertToInt8(request[currentIndex:currentIndex + 1]) - 1)
		currentIndex++

		topicName := request[currentIndex:currentIndex + topicNameLength]
		topicMetadata := CreateTopicMetadata(topicNameLength, string(topicName))
		topicMetadataSlice = append(topicMetadataSlice, topicMetadata)

		// move past the length
		currentIndex += topicNameLength
		
		// skip the buffer
		currentIndex++

		//processed the topic so decrease count
		topicsCount--
	}
	return topicMetadataSlice, currentIndex
}

func CreateTopicMetadata(topicNameLength int16, topicName string) TopicMetadata {
	var topicMetadata TopicMetadata
	topicMetadata.TopicNameLength = topicNameLength
	topicMetadata.TopicName = topicName
	return topicMetadata
}
