package internal

import (
	"bytes"
	"fmt"
)

func SerializeMap(data map[string]string, buf *bytes.Buffer) error {
	if data == nil {
        return fmt.Errorf("Map cannot be nil")
    }
    err := SerializeUint(uint64(len(data)), buf)
	if err != nil {
		return fmt.Errorf("error in SerializeMap: %w", err)
	}
	for key, value := range data {
		err := SerializeString(key, buf)
		if err != nil {
			return fmt.Errorf("error in SerializeMap: %w", err)
		}

		err = SerializeString(value, buf)
		if err != nil {
			return fmt.Errorf("error in SerializeMap: %w", err)
		}
	}
	return nil
}

func DeserializeMap(data *bytes.Buffer) (map[string]string, error) {
	var result map[string]string

	countEntry, err := DeserializeUint(data)
	if err != nil {
		return nil, fmt.Errorf("error in DeserializeMap: %w", err)
	}

	result = make(map[string]string, countEntry)

	for i := uint64(0); i < countEntry; i++ {
		key, err := DeserializeString(data)
		if err != nil {
			return nil, fmt.Errorf("error in DeserializeMap: %w", err)
		}
		value, err := DeserializeString(data)
		if err != nil {
			return nil, fmt.Errorf("error in DeserializeMap: %w", err)
		}
		result[key] = value
	}
	return result, nil
}
