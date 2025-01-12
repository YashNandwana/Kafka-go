package apiVersions

type ApiVersionsResponse struct {
	apiVersion int16
	minVersion int16
	maxVersion int16
	tagBuffer  int8
}

func (avr* ApiVersionsResponse) resposneGenerator() []byte {
	apiVersionResponse := []byte{
								0, byte(avr.apiVersion),
								0, byte(avr.minVersion),
								0, byte(avr.maxVersion),
								byte(avr.tagBuffer)}	
	return apiVersionResponse
}

func GenerateApiResponse(apiVersion, minVersion, maxVersion int16, tagBuffer int8) []byte {
	apiVersionResponse := GenerateApiPayload(apiVersion, minVersion, maxVersion, tagBuffer)
	response := apiVersionResponse.resposneGenerator()
	return response
}

func GenerateApiPayload(apiVersion, minVersion, maxVersion int16, tagBuffer int8) ApiVersionsResponse{
	var apiVersionResponse ApiVersionsResponse 
	apiVersionResponse.apiVersion = apiVersion
	apiVersionResponse.minVersion = minVersion
	apiVersionResponse.maxVersion = maxVersion
	apiVersionResponse.tagBuffer = tagBuffer

	return apiVersionResponse
}