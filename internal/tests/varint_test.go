package internal

import (
	"bytes"
	"math"
	"testing"
    "binary-serialization/internal"
	"github.com/stretchr/testify/require"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeVarint(t *testing.T) {
	t.Run("Unsigned zero", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := uint64(0)
		err := internal.SerializeUint(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeUint(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Max value", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
        input := uint64(math.MaxUint64)
		err := internal.SerializeUint(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeUint(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Signed zero", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := int64(0)
		err := internal.SerializeInt(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeInt(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

    t.Run("Signed min", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := int64(math.MinInt64)
		err := internal.SerializeInt(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeInt(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})
    
    t.Run("Signed max", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := int64(math.MaxInt64)
		err := internal.SerializeInt(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeInt(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})
    
    t.Run("Signed positive", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := int64(612132)
		err := internal.SerializeInt(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeInt(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})


    t.Run("Signed negative", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := int64(-450784)
		err := internal.SerializeInt(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeInt(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

    t.Run("Some unsigned value", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := uint64(10240124)
		err := internal.SerializeUint(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeUint(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

}

