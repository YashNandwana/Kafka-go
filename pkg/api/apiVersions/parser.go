package apiVersions


func ParseRequest(request []byte) ([]byte, error) {
	resp := make([]byte, 4)

	// parse correlation_id
	correlationId := request[8:12]
	
	// api version is 6 to 7 index
	requestedApiVersion := request[6:8]

	apiVersionErrorCode := GetApiVersionErrorCode(requestedApiVersion)
	
	resp = append(resp, correlationId...)
	resp = append(resp, apiVersionErrorCode...)
	
	resp = append(resp, 3) // number of api keys

	// generate resposne for all the apis
	apiVersionsResponse := GenerateApiResponse(18, 4, 4, 0)
	describeTopicPartitionsResponse := GenerateApiResponse(75, 0, 0, 0)
	
	resp = append(resp, apiVersionsResponse...)
	resp = append(resp, describeTopicPartitionsResponse...)

	suffixThrottleBuffer := []byte{0, 0, 0, 0,     // throttle
					  0, 			  // tag buffer
					}
	resp = append(resp, suffixThrottleBuffer...)
	resp = AddMessageSizeToResponse(resp)

	return resp, nil
	
}