package main

import (
	"binary-serialization/internal"
	"bufio"
	"fmt"
	"io"
	"os"
)


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

type UserEntry struct {
	key   string
	value string
	meta  map[string]string
}

func (u *UserEntry) Serialize(ser internal.Serializer) error {
    err := ser.SerializeString(u.key)
    check(err)
	err = ser.SerializeString(u.value)
    check(err)
	err = ser.SerializeStringMap(u.meta)
    check(err)
	return nil
}

func (u *UserEntry) Deserialize(deser internal.Deserializer) error {
	u.key, _ = deser.DeserializeString()
	u.value, _ = deser.DeserializeString()
	u.meta, _ = deser.DeserializeStringMap()
	return nil
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func writeFile(name string, data []byte) {
    f, err := os.Create(name)
    check(err)
    defer f.Close()


    w := bufio.NewWriter(f)
    _, err = w.Write(data)
    check(err)

    w.Flush()

}

func readFile(name string) []byte {
    
    f, err := os.Open(name)


    check(err)

    defer f.Close()

    reader := bufio.NewReader(f)
    buf := make([]byte, 1024)

    for {
        _, err := reader.Read(buf)

        if err != nil {
            if err != io.EOF {
                check(err)
            }
            break
        }
    }
    return buf
}


func serializationTest() {
	var (
		ser   internal.Serializer
		deser internal.Deserializer
	)


	ser = internal.NewBinarySerializer(2000)
    for i := 0; i < 1000; i++ {
        someStruct := UserEntry{
		    key:   fmt.Sprintf("Key_%d", i),
		    value: fmt.Sprintf("Value_%d", i),
		    meta: map[string]string{
			    "size": "100G",
			    "path": fmt.Sprintf("/user/home/folder_%d", i),
			    "cred": "rwxrwxrwx",
		    },
	    }
        someStruct.Serialize(ser)
    }


    bytes := ser.EndSerialize()
    fmt.Printf("cap = %d, len = %d\n", cap(bytes), len(bytes))
    
    // writeFile("/tmp/data", bytes)
    
    // readBytes := readFile("/tmp/data")

	deser = internal.NewBinaryDeserializer(bytes)


    for i := 0; i < 1000; i++ { 
	    newValue := UserEntry{}
        newValue.Deserialize(deser)
	    fmt.Printf("key: %s, value: %s, map:%v\n", newValue.key, newValue.value, newValue.meta)
	}
}



func main() {
    serializationTest()
}
