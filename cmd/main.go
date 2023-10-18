package main

import (
	"binary-serialization/internal"
	"fmt"
)

func Reverse[T any](input []T) []T {
	inputLen := len(input)
	output := make([]T, inputLen)

	for i, n := range input {
		j := inputLen - i - 1
		output[j] = n
	}
	return output
}

func debugSlice(buf []byte) {
	buf = Reverse(buf)
	fmt.Print("[ ")
	for _, byte := range buf {
		fmt.Printf("%08b ", byte)
	}
	fmt.Print("]\n")
}

type UserEntry struct {
	key   string
	value string
	meta  map[string]string
}

type Serializable interface {
	Serialize(ser internal.Serializer) ([]byte, error)
}

type Deserializable interface {
	Deserializable(deser internal.Deserializer) error
}

type StorageValue interface {
	Serializable
	Deserializable
}

func (u *UserEntry) Serialize(ser internal.Serializer) ([]byte, error) {
	ser.SerializeString(u.key)
	ser.SerializeString(u.value)
	ser.SerializeStringMap(u.meta)
	return ser.EndSerialize(), nil
}

func (u *UserEntry) Deserialize(deser internal.Deserializer) error {
	u.key, _ = deser.DeserializeString()
	u.value, _ = deser.DeserializeString()
	u.meta, _ = deser.DeserializeStringMap()
	return nil
}

func main() {
	var (
		ser   internal.Serializer
		deser internal.Deserializer
	)

	ser = internal.NewBinarySerializer(10)
	someStruct := UserEntry{
		key:   "SomeKey",
		value: "SomeValue",
		meta: map[string]string{
			"size": "100",
			"path": "/user/home/folder",
			"cred": "rwxrwxrwx",
		},
	}

	bytes, _ := someStruct.Serialize(ser)

	deser = internal.NewBinaryDeserializer(bytes)

	newValue := UserEntry{}
	newValue.Deserialize(deser)
	fmt.Printf("key: %s, value: %s, map:%v\n", newValue.key, newValue.value, newValue.meta)
}
