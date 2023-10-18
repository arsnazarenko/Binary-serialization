package main

import (
    bin "binary-serialization/internal"
//	"encoding/binary"
	"fmt"
	"math"

	"bytes"
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

// print varints in buf
func debugSlice(buf []byte) {
	buf = Reverse(buf)
	fmt.Print("[[ ")
	for _, byte := range buf {
		fmt.Printf("%08b ", byte)
	}
	fmt.Print("]\n")
}

func main() {

    
    

//    init := []string{"Hello", "world", "Роман", "ЮТФ-8 Рулит!!!"}
//    
//    for _, s := range init {
//        bin.SerializeString(s, buf)
//        fmt.Printf("cap: %d, len: %d\n", buf.Cap(), buf.Len())
//        res, _ := bin.DeserializeString(buf)
//        fmt.Printf("init: %s, result: %s\n", s, res)
//        fmt.Printf("cap: %d, len: %d\n", buf.Cap(), buf.Len())
//    }
    
    buf := bytes.NewBuffer(make([]byte, 0, 50))
    arr := []uint64{0, 1, 0xFFFF, math.MaxUint64}
    for _, v := range arr {
        fmt.Printf("cap: %d, len: %d\n", buf.Cap(), buf.Len())
        bin.SerializeUint(v, buf)
        fmt.Printf("cap: %d, len: %d\n", buf.Cap(), buf.Len())    
    }
    s := buf.Bytes()
    fmt.Printf("cap: %d, len: %d\n", buf.Cap(), buf.Len())    
    fmt.Printf("Slice cap: %d, slice len: %d\n", cap(s), len(s))
    _ = buf
    

    newBuf := bytes.NewBuffer(s)
    for _, v := range arr {
        fmt.Printf("cap: %d, len: %d\n", newBuf.Cap(), newBuf.Len())
        res, _ := bin.DeserializeUint(newBuf)
        fmt.Printf("value: %d, result: %d\n", v, res)
        fmt.Printf("cap: %d, len: %d\n", newBuf.Cap(), newBuf.Len())
    }
    fmt.Printf("cap: %d, len: %d\n", newBuf.Cap(), newBuf.Len())    
    
}
