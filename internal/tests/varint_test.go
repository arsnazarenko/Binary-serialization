package internal

import (
	"math"
	"testing"
    "binary-serialization/internal"
	"github.com/stretchr/testify/require"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeVarint(t *testing.T) {
	t.Run("Unsigned zero", func(t *testing.T) {

        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := uint64(0)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeUint(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeUint()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})

	t.Run("Unsigned max value", func(t *testing.T) {

        
        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := uint64(math.MaxUint64)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeUint(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeUint()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})

    t.Run("Some unsigned value", func(t *testing.T) {
   
     var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := uint64(10000001)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeUint(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeUint()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})



	t.Run("Signed zero", func(t *testing.T) {
       
        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := int64(0)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeInt(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeInt()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})

    t.Run("Signed min", func(t *testing.T) {
        
        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := int64(math.MinInt64)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeInt(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeInt()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})
    
    t.Run("Signed max", func(t *testing.T) {
        

        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := int64(math.MaxInt64)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeInt(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeInt()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})
    
    t.Run("Signed positive", func(t *testing.T) {

        
        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := int64(980765)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeInt(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeInt()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})


    t.Run("Signed negative", func(t *testing.T) {

        
        var (
			ser   *internal.BinarySerializer
			deser *internal.BinaryDeserializer
		)

		input := int64(-980765)
        
        ser = internal.NewBinarySerializer( /*capacity*/ 0)
		err := ser.SerializeInt(input)
		require.NoError(t, err)

		bytes := ser.EndSerialize()
		deser = internal.NewBinaryDeserializer(bytes)

		actual, err := deser.DeserializeInt()
		require.NoError(t, err)

		require.Equal(t, input, actual)

	})
}

