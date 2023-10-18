package internal

import (
	"bytes"
	"math"
	"testing"
    "binary-serialization/internal"
	"github.com/stretchr/testify/require"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeFloat(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := 0.0
		err := internal.SerializeFloat(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeFloat(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Max value", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := math.MaxFloat64
		err := internal.SerializeFloat(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeFloat(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Min value", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := -math.MaxFloat64
		err := internal.SerializeFloat(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeFloat(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

    t.Run("Some value", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := 1123.1239871793
		err := internal.SerializeFloat(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeFloat(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})
}
