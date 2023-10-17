package internal

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeString(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		input := "hello"
		bytes, err := SerializeString(input)
		require.NoError(t, err)

		actual, offset, err := DeserializeString(bytes)
		require.NoError(t, err)
		fmt.Println(offset)
		require.Equal(t, input, actual)
	})

	t.Run("How are you baby?", func(t *testing.T) {
		input := "How are you baby?"
		bytes, err := SerializeString(input)
		require.NoError(t, err)

		actual, offset, err := DeserializeString(bytes)
		require.NoError(t, err)
		fmt.Println(offset)
		require.Equal(t, input, actual)
	})

	t.Run("", func(t *testing.T) {
		input := ""
		bytes, err := SerializeString(input)
		require.NoError(t, err)

		actual, offset, err := DeserializeString(bytes)
		require.NoError(t, err)
		fmt.Println(offset)
		require.Equal(t, input, actual)
	})

}
