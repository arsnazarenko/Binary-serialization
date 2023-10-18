package internal

import (
	"math"
	"testing"
    "binary-serialization/internal"
	"github.com/stretchr/testify/require"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeFloat(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {

        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := 0.0
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeFloat(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeFloat()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})

	t.Run("Max value", func(t *testing.T) {

 
        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := math.MaxFloat64
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeFloat(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeFloat()
		require.NoError(t, err)

		require.Equal(t, input, actual)       

	})

	t.Run("Min value", func(t *testing.T) {

        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := -math.MaxFloat64
        
        ser = internal.NewBinarySerializer( /*capacity*/ 50)
		err := ser.SerializeFloat(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeFloat()
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

    t.Run("Some value", func(t *testing.T) {

        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := 1123.1239871793
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeFloat(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeFloat()
		require.NoError(t, err)

		require.Equal(t, input, actual)


	})
}
