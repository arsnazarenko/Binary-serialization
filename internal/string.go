package internal

import (
	"bytes"
	"fmt"
)

// Binary schema for string (for payload use UTF-8)
// [type] [payload_uint] [payload]...

const (
	StringTypeByte byte = 0x3 // StringTypeByte like 00000011
)

func SerializeString(data string, buf *bytes.Buffer) error {

	err := SerializeUint(uint64(len(data)), buf)
	if err != nil {
		return fmt.Errorf("error in SerializeString: %w", err)
	}
	buf.Write([]byte(data))
	return nil
}

func DeserializeString(data *bytes.Buffer) (string, error) {
    countBytes, err := DeserializeUint(data)
	if err != nil {
		return "", fmt.Errorf("error in DeserializeString: %w", err)
	}
	strBytes := make([]byte, countBytes, countBytes)
	_, err = data.Read(strBytes)
	if err != nil {
		return "", fmt.Errorf("erro in _deserializeString: %w", err)
	}
	str := string(strBytes)
	return str, nil
}

