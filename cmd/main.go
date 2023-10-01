package main

import (
	"fmt"
	"math"
	"math/bits"
)

type Serializable interface {
	Serialize(ser Serializer) ([]byte, error)
}

type Serializer interface {
	SerializeInt8(value int8) error
	SerializeInt16(value int16) error
	SerializeInt32(value int32) error
	SerializeInt64(value int64) error
	SerializeUint8(value uint8) error
	SerializeUint16(value uint16) error
	SerializeUint32(value uint32) error
	SerializeUint64(value uint64) error
	SerializeString(value string) error
	SerializeSlise(value []any) error
	End() ([]byte, error)
}

// Example: It's our binary format serializator, for json you can create JsonSerializator and impl Serializator interface
type BinarySerializer struct {
	buf []byte
}

func New(capacity uint) *BinarySerializer {
	return &BinarySerializer{
		buf: make([]byte, 0, capacity),
	}
}

/*
	In SerializeXXX methods data type perform to to binary array and
	write into internal buf of conrete type of Serializator.
	In the end of serialization you can call End method and get
	slise of bytes.

	For user-defined structs you should implement Serializable interface

	Example:

	type User struct {
	    age uint8,
	    name string,
	}

	func (u *User) Serialize(ser Serializator) []byte, error {
	    ser.SerializeUint8(u.age) // append bytes to buf
	    ser.SerializeString(u.name) //append bytes to buf
	    return (ser.End(), nil) // return buf
	}
*/

func IntegerToVarint(value uint64, buf []byte) error {
	return nil
}

func (b *BinarySerializer) SerializeInt8(value int8) error {
	return nil
}

func (b *BinarySerializer) SerializeInt16(value int16) error {
	return nil
}

func (b *BinarySerializer) SerializeInt32(value int32) error {
	return nil
}

func (b *BinarySerializer) SerializeInt64(value int64) error {
	return nil
}

func (b *BinarySerializer) SerializeUint8(value uint8) error {
	return nil
}

func (b *BinarySerializer) SerializeUint16(value uint16) error {
	return nil
}

func (b *BinarySerializer) SerializeUint32(value uint32) error {
	return nil
}

func (b *BinarySerializer) SerializeUint64(value uint64) error {
	return nil
}

func (b *BinarySerializer) SerializeString(value string) error {
	return nil
}

func (b *BinarySerializer) SerializeSlise(value []any) error {
	return nil
}

func (b *BinarySerializer) End() ([]byte, error) {
	return nil, nil
}

func UnsignedToVarint(buf []byte, val uint64) ([]byte, error) {

	// todo: append prefix with type of data and some additional info (count for arrays and etc.)

	bitsLen := bits.Len64(val)
	remainder := bitsLen % 7
	bytesLen := bitsLen / 7
	if remainder > 0 {
		bytesLen += 1
	}
	for i := 0; i < bytesLen; i++ {
		curByte := byte((val >> (7 * i)) & (0x7f))
		if i < bytesLen-1 {
			curByte |= 0x80
		}
		buf = append(buf, curByte)
	}
	return buf, nil
}


func SignedToVarint(buf []byte, val int64) error {
	return nil
}

// print varints in buf
func debugSlice(buf []byte) {
    for _, byte := range buf {
        fmt.Printf("%08b ", byte)
        if (byte & 0x80) == 0 {
            fmt.Print("\n")
        }
	}
}

func main() {
	buf := make([]byte, 0, 10)
	buf, _ = UnsignedToVarint(buf, 1)
	buf, _ = UnsignedToVarint(buf, 8)
	buf, _ = UnsignedToVarint(buf, 16)
	buf, _ = UnsignedToVarint(buf, 32)
	buf, _ = UnsignedToVarint(buf, 128) // in to bytes
	buf, _ = UnsignedToVarint(buf, 512) // in to bytes
	buf, _ = UnsignedToVarint(buf, 1024) // in to bytes
	buf, _ = UnsignedToVarint(buf, 1024) // in to bytes
	buf, _ = UnsignedToVarint(buf, math.MaxUint64) // in to bytes
	debugSlice(buf)
}
