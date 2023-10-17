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

func _serializeString(data string) ([]byte, error) {
	var buf bytes.Buffer

	bytesLen, err := SerializeUint(uint64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("error in _serializeString: %w", err)
	}
	buf.Write(bytesLen[1:])
	buf.Write([]byte(data))

	return buf.Bytes(), nil
}

// ok
func _deserializeString(data []byte) (string, uint64, error) {
	var (
		buf    bytes.Buffer
		offset uint64
	)

	countBytes, uintOffset, err := _deserializeUint(data)
	if err != nil {
		return "", 0, fmt.Errorf("error in _deserializeString: %w", err)
	}
	offset += uintOffset

	buf.Write(data[offset : offset+countBytes])
	offset += countBytes

	return buf.String(), offset, nil
}

// SerializeString serializes string to []byte
func SerializeString(data string) ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteByte(StringTypeByte)
	payload, err := _serializeString(data)
	if err != nil {
		return nil, fmt.Errorf("error in SerializeString: %w", err)
	}
	buf.Write(payload)

	return buf.Bytes(), nil
}

// DeserializeString serializes []byte to string
func DeserializeString(data []byte) (string, uint64, error) {
	if data[0] != StringTypeByte {
		return "", 0, fmt.Errorf("invalid row data")
	}

	result, offset, err := _deserializeString(data[1:])
	if err != nil {
		return "", 0, err
	}

	return result, offset + 1, nil
}
