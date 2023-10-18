package internal

import (
	"binary-serialization/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeString(t *testing.T) {
	t.Run("Long string", func(t *testing.T) {

		var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := " How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend? How are you, my friend?"

		ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeString(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeString()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})

	t.Run("Little string", func(t *testing.T) {

		var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := "Hello, world!"
		ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeString(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeString()
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Zero string", func(t *testing.T) {

		var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := ""
		ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeString(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeString()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})

	t.Run("UTF-8 string", func(t *testing.T) {

		var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := "Проверка работы русских символов"
		ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeString(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeString()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})

}
