package internal

import (
	"binary-serialization/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeMap(t *testing.T) {
	t.Run("Map with one element", func(t *testing.T) {
		var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := map[string]string{
			"Roman": "Golang Developer",
		}

		ser = internal.NewBinarySerializer( /*capacity*/ 50)
		err := ser.SerializeStringMap(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeStringMap()
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Multiply elements map", func(t *testing.T) {
		var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := map[string]string{
			"Arseniy": "Rust",
			"Roman":   "Golang",
			"Maria":   "Kotlin",
			"Artem":   "Python",
			"Igor":    "OCaml",
		}

		ser = internal.NewBinarySerializer( /*capacity*/ 50)
		err := ser.SerializeStringMap(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeStringMap()
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Empty map", func(t *testing.T) {

		var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := map[string]string{}

		ser = internal.NewBinarySerializer( /*capacity*/ 50)
		err := ser.SerializeStringMap(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeStringMap()
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

}
