package internal

import (
	"bytes"
	"testing"
    "binary-serialization/internal"
	"github.com/stretchr/testify/require"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeString(t *testing.T) {
	t.Run("Little string", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := "hello"
		err := internal.SerializeString(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeString(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Long string", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := " How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby? How are you baby?"
		err := internal.SerializeString(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeString(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Zero string", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := ""
		err := internal.SerializeString(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeString(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

}
