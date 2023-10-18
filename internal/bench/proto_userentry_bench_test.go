package bench

import (
	"binary-serialization/internal/protobuf"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkSerializableProto(b *testing.B) {
	cases := []struct {
		Name  string
		Key   string
		Value string
		Meta  map[string]string
	}{
		{Name: "Value len-1", Key: "key1", Value: "a", Meta: map[string]string{"meta-key": "meta-value"}},
		{Name: "Value len-100", Key: "key2", Value: strings.Repeat("a", 100), Meta: map[string]string{"meta-key": "meta-value"}},
		{Name: "Value len-10000", Key: "key3", Value: strings.Repeat("a", 10000), Meta: map[string]string{"meta-key": "meta-value"}},
		{Name: "Map len-1", Key: "key3", Value: "a", Meta: genMap(1)},
		{Name: "Map len-100", Key: "key3", Value: "a", Meta: genMap(100)},
		{Name: "Map len-10000", Key: "key3", Value: "a", Meta: genMap(10000)},
	}

	for _, tc := range cases {
		b.Run(tc.Name, func(b *testing.B) {
			userKeyValue := &protobuf.UserEntry{
				Key:   tc.Key,
				Value: tc.Value,
				Meta:  tc.Meta,
			}

			for i := 0; i < b.N; i++ {
				_, err := proto.Marshal(userKeyValue)
				if err != nil {
					log.Fatal("Ошибка при сериализации:", err)
				}
			}
		})
	}
}

func BenchmarkDeserializableProto(b *testing.B) {
	tcBin := []struct {
		Name string
		Data []byte
	}{
		{Name: "Value len-1"},
		{Name: "Value len-100"},
		{Name: "Value len-10000"},
		{Name: "Map len-1"},
		{Name: "Map len-100"},
		{Name: "Map len-10000"},
	}
	preTestic := []*protobuf.UserEntry{
		{Key: "key1", Value: "a", Meta: map[string]string{"meta-key": "meta-value"}},
		{Key: "key2", Value: strings.Repeat("a", 100), Meta: map[string]string{"meta-key": "meta-value"}},
		{Key: "key3", Value: strings.Repeat("a", 10000), Meta: map[string]string{"meta-key": "meta-value"}},
		{Key: "key3", Value: "a", Meta: genMap(1)},
		{Key: "key3", Value: "a", Meta: genMap(100)},
		{Key: "key3", Value: "a", Meta: genMap(10000)},
	}
	for i, v := range preTestic {
		data, err := proto.Marshal(v)
		if err != nil {
			log.Fatal("Ошибка при сериализации:", err)
		}
		tcBin[i].Data = data
	}

	for _, tc := range tcBin {
		b.Run(tc.Name, func(b *testing.B) {
			newUserEntry := &protobuf.UserEntry{}
			for i := 0; i < b.N; i++ {
				err := proto.Unmarshal(tc.Data, newUserEntry)
				if err != nil {
					log.Fatal("Ошибка при сериализации:", err)
				}
			}
		})
	}

}

func genMap(countElem int) map[string]string {
	m := make(map[string]string, countElem)

	for i := 0; i < countElem; i++ {
		m[strconv.Itoa(i)+"key"] = fmt.Sprintf("value-%d", i)
	}

	return m
}
