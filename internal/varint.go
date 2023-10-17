package internal

import "math/bits"
import "fmt"

const (
	VarintType  byte = 0x1 // protobuf Varint type for unsigned numbers
	SignIntType byte = 0x2 // protobuf sint32/sint64 type for signed numbers with zig-zag coding
)

func _serializeUintInto(value uint64, buf []byte) ([]byte, error) {
	bitsLen := bits.Len64(value)
	if bitsLen == 0 {
		buf = append(buf, 0)
		return buf, nil
	}
	bytesLen, remainder := bitsLen/7, bitsLen%7
	if remainder > 0 {
		bytesLen++
	}
	for i := 0; i < bytesLen; i++ {
		buf = append(buf, byte((value>>(7*i))&0x7f|0x80))
	}
	buf[len(buf)-1] &= 0x7f
	return buf, nil

}

func _deserializeUint(data []byte) (uint64, uint64, error) {
	res := uint64(0)
	i := 0
	for {
		res |= uint64((data[i] & 0x7F)) << (7 * i)
		if (data[i] & 0x80) == 0 {
			break
		}
		i++
	}
	return res, uint64(i + 1), nil
}

func SerializeUint(value uint64) ([]byte, error) {
	buf := make([]byte, 0, 10)
	buf = append(buf, VarintType)
	return _serializeUintInto(value, buf)
}

func DeserializeUint(data []byte) (uint64, uint64, error) {
	if len(data) < 2 { // tag and one byte is minimal size of varint
		return 0, 0, fmt.Errorf("Invalid row data")
	}
	if data[0] != VarintType {
		return 0, 0, fmt.Errorf("Buffer contains the wrong data type")
	}
    value, offset, err := _deserializeUint(data[1:])
    offset++
    return value, offset, err
}

func SerializeInt(value int64) ([]byte, error) {
	buf := make([]byte, 0, 10)
	buf = append(buf, SignIntType)
	unsigned := uint64((value << 1) ^ (value >> (64 - 1)))
	return _serializeUintInto(unsigned, buf)
}

func DeserializeInt(data []byte) (int64, uint64, error) {
	if len(data) < 2 { // tag and one byte is minimal size of varint
		return 0, 0, fmt.Errorf("Invalid row data")
	}
	if data[0] != SignIntType {
		return 0, 0, fmt.Errorf("Buffer contains the wrong data type")
	}
	unsigned, offset, err := _deserializeUint(data[1:])
	if err != nil {
		return 0, 0, err
	}
    offset++
	result := int64((unsigned >> 1) ^ -(unsigned & 0x1))
	return result, offset, nil
}
