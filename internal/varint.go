package internal

import (
	"bytes"
	"math/bits"
)


func SerializeUint(value uint64, buf *bytes.Buffer) error {
	bitsLen := bits.Len64(value)
	if bitsLen == 0 {
		buf.WriteByte(0)
        return nil
	}
	bytesLen, remainder := bitsLen/7, bitsLen%7
	if remainder > 0 {
		bytesLen++
	}
	for i := 0; i < bytesLen; i++ {
        curByte := byte((value>>(7*i))&0x7f|0x80)
        if i == (bytesLen - 1) {
            curByte &= 0x7f
        }
		buf.WriteByte(curByte)
	}
	return nil
}

func DeserializeUint(data *bytes.Buffer) (uint64, error) {
	res := uint64(0)
	i := 0
    for  {
        curByte, err := data.ReadByte();
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


func SerializeInt(value int64, buf *bytes.Buffer) error {
	unsigned := uint64((value << 1) ^ (value >> (64 - 1)))
	return SerializeUint(unsigned, buf)
}

func DeserializeInt(data *bytes.Buffer) (int64, error) {
	unsigned, err := DeserializeUint(data)
	if err != nil {
		return 0, err
	}
	result := int64((unsigned >> 1) ^ -(unsigned & 0x1))
	return result, nil
}
