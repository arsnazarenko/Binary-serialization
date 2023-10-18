package internal

import (
	"bytes"
	"encoding/binary"
	"math"
	"fmt"
)


func SerializeFloat(value float64, buf *bytes.Buffer) error {
    numSlice := make([]byte, 0, 8) 
    numSlice = binary.LittleEndian.AppendUint64(numSlice, math.Float64bits(value))
    _, err := buf.Write(numSlice)
    if err != nil {
        return fmt.Errorf("error in SerializeFloat: %w", err)
    }
    return nil
}

func DeserializeFloat(data *bytes.Buffer) (float64, error) {
    value := binary.LittleEndian.Uint64(data.Next(8))
    res := math.Float64frombits(value)
    return res, nil
}

