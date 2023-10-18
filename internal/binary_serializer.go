package internal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/bits"
)

type BinarySerializer struct {
	buf *bytes.Buffer
}

func NewBinarySerializer(capacity int) *BinarySerializer {
	return &BinarySerializer{
        buf: bytes.NewBuffer(make([]byte, 0, capacity)),
    }
}

func (bs *BinarySerializer) SerializeUint(value uint64) error {
	bitsLen := bits.Len64(value)
	if bitsLen == 0 {
		bs.buf.WriteByte(0)
		return nil
	}
	bytesLen, remainder := bitsLen/7, bitsLen%7
	if remainder > 0 {
		bytesLen++
	}
	for i := 0; i < bytesLen; i++ {
		curByte := byte((value>>(7*i))&0x7f | 0x80)
		if i == (bytesLen - 1) {
			curByte &= 0x7f
		}
		bs.buf.WriteByte(curByte)
	}
	return nil
}

func (bs *BinarySerializer) SerializeInt(value int64) error {
	unsigned := uint64((value << 1) ^ (value >> (64 - 1)))
	return bs.SerializeUint(unsigned)
}

func (bs *BinarySerializer) SerializeString(value string) error {
	err := bs.SerializeUint(uint64(len(value)))
	if err != nil {
		return fmt.Errorf("error in SerializeString: %w", err)
	}
	bs.buf.Write([]byte(value))
	return nil
}

func (bs *BinarySerializer) SerializeFloat(value float64) error {
	numSlice := make([]byte, 0, 8)
	numSlice = binary.LittleEndian.AppendUint64(numSlice, math.Float64bits(value))
	_, err := bs.buf.Write(numSlice)
	if err != nil {
		return fmt.Errorf("error in SerializeFloat: %w", err)
	}
	return nil
}

func (bs *BinarySerializer) SerializeStringMap(value map[string]string) error {
	if value == nil {
		return fmt.Errorf("Map cannot be nil")
	}
	err := bs.SerializeUint(uint64(len(value)))
	if err != nil {
		return fmt.Errorf("error in SerializeMap: %w", err)
	}
	for key, value := range value {
		err := bs.SerializeString(key)
		if err != nil {
			return fmt.Errorf("error in SerializeMap: %w", err)
		}

		err = bs.SerializeString(value)
		if err != nil {
			return fmt.Errorf("error in SerializeMap: %w", err)
		}
	}
	return nil
}

func (bs *BinarySerializer) EndSerialize() []byte {
	src := bs.buf.Bytes()
	dst := make([]byte, len(src))
	copy(dst, src)
	bs.buf.Reset()
	return dst
}



type BinaryDeserializer struct {
    buf *bytes.Buffer
}


func NewBinaryDeserializer(from []byte) *BinaryDeserializer {
    return &BinaryDeserializer{
        buf: bytes.NewBuffer(from),
    }
}

func (bd *BinaryDeserializer) DeserializeUint() (uint64, error) {
	res := uint64(0)
	i := 0
	for {
		curByte, err := bd.buf.ReadByte()
		if err != nil {
			return 0, nil
		}
		res |= uint64((curByte & 0x7F)) << (7 * i)
		if (curByte & 0x80) == 0 {
			break
		}
		i++
	}
	return res, nil

}

func (bd *BinaryDeserializer) DeserializeInt() (int64, error) {
	unsigned, err := bd.DeserializeUint()
	if err != nil {
		return 0, err
	}
	result := int64((unsigned >> 1) ^ -(unsigned & 0x1))
	return result, nil
}

func (bd *BinaryDeserializer) DeserializeString() (string, error) {
	countBytes, err := bd.DeserializeUint()
	if err != nil {
		return "", fmt.Errorf("error in DeserializeString: %w", err)
	}
	strBytes := make([]byte, countBytes, countBytes)
	_, err = bd.buf.Read(strBytes)
	if err != nil {
		return "", fmt.Errorf("error in _deserializeString: %w", err)
	}
	str := string(strBytes)
	return str, nil
}

func (bd *BinaryDeserializer) DeserializeFloat() (float64, error) {
	value := binary.LittleEndian.Uint64(bd.buf.Next(8))
	res := math.Float64frombits(value)
	return res, nil
}

func (bd *BinaryDeserializer) DeserializeStringMap() (map[string]string, error) {
	var result map[string]string
	countEntry, err := bd.DeserializeUint()
	if err != nil {
		return nil, fmt.Errorf("error in DeserializeMap: %w", err)
	}
	result = make(map[string]string, countEntry)

	for i := uint64(0); i < countEntry; i++ {
		key, err := bd.DeserializeString()
		if err != nil {
			return nil, fmt.Errorf("error in DeserializeMap: %w", err)
		}
		value, err := bd.DeserializeString()
		if err != nil {
			return nil, fmt.Errorf("error in DeserializeMap: %w", err)
		}
		result[key] = value
	}
	return result, nil
}

func (bd *BinaryDeserializer) Reset() error {
	internal := bd.buf.Bytes()
	bd.buf = bytes.NewBuffer(internal)
	return nil
}
