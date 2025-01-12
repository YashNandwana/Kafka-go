package apiVersions

import "encoding/binary"

func AddMessageSizeToResponse(response []byte) []byte {
	response[3] = byte(len(response) - 4)
	return response
}

func GetApiVersionErrorCode(api_version []byte) []byte {
	var (
		version = binary.BigEndian.Uint16(api_version)
		api_version_error_code []byte
	)
	switch version{
	case 0, 1, 2, 3, 4:
		api_version_error_code = []byte{0, 0}
	default:
		api_version_error_code = []byte{0, 35}
	}
	return api_version_error_code
}