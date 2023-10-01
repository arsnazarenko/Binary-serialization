package main

import "fmt"

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

func main() {

	fmt.Printf("Hello, serialization!\n")
}
