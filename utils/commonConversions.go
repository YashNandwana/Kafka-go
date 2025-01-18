package utils

import (
	"encoding/binary"
)

func ConvertToInt16(byteSlice []byte) int16 {
	return int16(binary.BigEndian.Uint16(byteSlice))
}

func ConvertToInt8(byteSlice []byte) int8 {
    return int8(byteSlice[0])
}