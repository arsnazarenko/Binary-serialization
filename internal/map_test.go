package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeMap(t *testing.T) {
	t.Run("one elem", func(t *testing.T) {
		input := map[string]string{
			"roma": "lox",
		}
		bytes, err := SerializeMap(input)
		require.NoError(t, err)

		actual, _, err := DeserializeMap(bytes)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("two elem", func(t *testing.T) {
		input := map[string]string{
			"roma": "lox",
			"ars":  "kros",
		}
		bytes, err := SerializeMap(input)
		require.NoError(t, err)

		actual, _, err := DeserializeMap(bytes)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("empty mapa", func(t *testing.T) {
		input := map[string]string{}
		bytes, err := SerializeMap(input)
		require.NoError(t, err)

		actual, _, err := DeserializeMap(bytes)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

}
