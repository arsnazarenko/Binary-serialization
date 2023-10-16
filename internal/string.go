package internal

import (
	"bytes"
	"fmt"
)

// Binary schema for string (for payload use UTF-8)
// [type] [len] [payload]...

const (
	StringType    byte = 0x3  // StringType like 00000011
	MaxCountBytes byte = 0xFF // MaxCountBytes like 11111111
)

// SerializeString serializes string to []byte
func SerializeString(data string) ([]byte, error) {
	if len(data) > int(MaxCountBytes) {
		return nil, fmt.Errorf("too long string")
	}

	result := make([]byte, 0, 2+len(data))
	result = append(result, StringType, byte(len(data)))
	result = append(result, []byte(data)...)

	return result, nil
}

// DeserializeString serializes []byte to string
func DeserializeString(data []byte) (string, error) {
	var (
		result            bytes.Buffer
		countBytesPayload = int(data[1])
	)

	_, err := result.Write(data[2 : 2+countBytesPayload])
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
