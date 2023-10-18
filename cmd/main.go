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


type UserKeyValue {
    key string
    value string
    meta map[string]string
}

type Serializable interface {
    Serialize(ser internal.Serializer) ([]byte, error)
}

type Deserializable interface {
    Deserializable[](deser internal.Derializer) (error)
}   

func main() {
    
}
