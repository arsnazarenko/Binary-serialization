package main

import (
	bin "binary-serialization/internal"
	"fmt"
	"math"
	// "math"
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
	
    array := []int64{0, -10, 10, -65564, math.MinInt64, math.MaxInt64}
	for _, value := range array {
        outBuf, _ := bin.SerializeInt(value)
        res, offset, _ := bin.DeserializeInt(outBuf)
        fmt.Printf("original: %d\n  result: %d\n", value, res)
        fmt.Printf("offset: %d, ", offset)
        debugSlice(outBuf)
        fmt.Println()
	}
}
