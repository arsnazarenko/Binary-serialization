package internal

import (
	"bytes"
	"fmt"
)

const (
	MapTypeByte byte = 0x3
)

func _serializeMap(data map[string]string) ([]byte, error) {
	var buf bytes.Buffer

	binCountEntry, err := SerializeUint(uint64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("error in _serializeMap: %w", err)
	}
	buf.Write(binCountEntry[1:])

	for key, value := range data {
		binKey, err := _serializeString(key)
		if err != nil {
			return nil, fmt.Errorf("error in _serializeMap: %w", err)
		}
		buf.Write(binKey)

		binValue, err := _serializeString(value)
		if err != nil {
			return nil, fmt.Errorf("error in _serializeMap: %w", err)
		}
		buf.Write(binValue)
	}

	return buf.Bytes(), nil
}

func _deserializeMap(data []byte) (map[string]string, uint64, error) {
	var (
		result map[string]string
		offset uint64
	)

	countEntry, uintOffset, err := _deserializeUintMOCK(data)
	if err != nil {
		return nil, -1, fmt.Errorf("error in _deserializeMap: %w", err)
	}
	offset += uintOffset

	result = make(map[string]string, countEntry)

	var i uint64
	for ; i < countEntry; i++ {
		key, strKeyOffset, err := _deserializeString(data[offset:])
		if err != nil {
			return nil, -1, fmt.Errorf("error in _deserializeMap: %w", err)
		}
		offset += strKeyOffset

		value, strValueOffset, err := _deserializeString(data[offset:])
		if err != nil {
			return nil, -1, fmt.Errorf("error in _deserializeMap: %w", err)
		}
		offset += strValueOffset

		result[key] = value
	}

	return result, offset, nil
}

func SerializeMap(data map[string]string) ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteByte(MapTypeByte)    // map type
	buf.WriteByte(StringTypeByte) // string type for keys
	buf.WriteByte(StringTypeByte) // string type for values
	payload, err := _serializeMap(data)
	if err != nil {
		return nil, fmt.Errorf("error in SerializeMap: %w", err)
	}
	buf.Write(payload)

	return buf.Bytes(), nil
}

func DeserializeMap(data []byte) (map[string]string, uint64, error) {
	if data[0] != MapTypeByte || data[1] != StringTypeByte || data[2] != StringTypeByte {
		return nil, -1, fmt.Errorf("invalid row data")
	}

	result, offset, err := _deserializeMap(data[3:])
	if err != nil {
		return nil, -1, err
	}
	return result, offset + 3, nil
}
