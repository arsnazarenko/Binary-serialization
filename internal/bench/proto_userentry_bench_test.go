package bench

import (
	"binary-serialization/internal/protobuf"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

func BenchmarkSerializableProto(b *testing.B) {
	userKeyValue := &protobuf.UserEntry{
		Key:   "some-key",
		Value: "some-value",
		Meta:  map[string]string{"meta-key": "meta-value"},
	}
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(userKeyValue)
		if err != nil {
			log.Fatal("Ошибка при сериализации:", err)
		}
	}
}

func BenchmarkDeserializableProto(b *testing.B) {
	userKeyValue := &protobuf.UserEntry{
		Key:   "some-key",
		Value: "some-value",
		Meta:  map[string]string{"meta-key": "meta-value"},
	}
	data, err := proto.Marshal(userKeyValue)
	if err != nil {
		log.Fatal("Ошибка при сериализации:", err)
	}

	newUserEntry := &protobuf.UserEntry{}
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(data, newUserEntry)
		if err != nil {
			log.Fatal("Ошибка при сериализации:", err)
		}
	}
}
