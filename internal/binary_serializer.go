package internal

import (
	"binary-serialization/internal"
	"bytes"
)


type BinarySerializer struct {
    buf bytes.Buffer
}

func NewBinarySerializerFrom(buf []byte) *BinarySerializer {
    return &BinarySerializer{
    	buf: *bytes.NewBuffer(buf),
    }
}

func NewBinarySerializer() *BinarySerializer {
    return &BinarySerializer{
        buf: bytes.Buffer{},
    }
}

func (b *BinarySerializer) SerializeUint(value uint64) error {
    
}

func (b *BinarySerializer) SerializeInt(value int64) error {

}

func (b *BinarySerializer) SerializeString(value string) error {

}

func (b *BinarySerializer) SerializeFloat(value float64) error {

}

func (b *BinarySerializer) SerializeStringMap(value map[string]string) error {
    
}


func (b *BinarySerializer) EndSerialize() []byte { 
    src := b.buf.Bytes()
    dst := make([]byte, len(src))
    copy(dst, src)
    b.buf.Reset()
    return dst
}

func (b *BinarySerializer) DeserializeUint() (uint64, error) {}

func (b *BinarySerializer) DeserializeInt() (int64, error) {}

func (b *BinarySerializer) DeserializeString() (string, error) {}

func (b *BinarySerializer) DeserializeFloat() (float64, error) {}

func (b *BinarySerializer) DeserializeStringMap() (map[string]string, error) {}

func (b *BinarySerializer) Reset() error {
    internal := b.buf.Bytes()
    b.buf = *bytes.NewBuffer(internal)
    return nil
}


