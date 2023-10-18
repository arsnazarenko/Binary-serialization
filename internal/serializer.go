package internal

type Serializer interface {
	SerializeUint(value uint64) error
	SerializeInt(value int64) error
	SerializeString(value string) error
	SerializeFloat(value float64) error
	SerializeStringMap(value map[string]string) error
	EndSerialize() []byte
}

type Deserializer interface {
	DeserializeUint() (uint64, error)
	DeserializeInt() (int64, error)
	DeserializeString() (string, error)
	DeserializeFloat() (float64, error)
	DeserializeStringMap() (map[string]string, error)
	Reset() error
}
