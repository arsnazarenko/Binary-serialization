package internal

import (
	"bytes"
	"testing"
    "binary-serialization/internal"
	"github.com/stretchr/testify/require"
)

// Check Serialize - Deserialize
func TestSerializeDeserializeMap(t *testing.T) {
	t.Run("Map with one element", func(t *testing.T) {
            
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := map[string]string{
			"Roman": "Golang Developer",
		}
		err := internal.SerializeMap(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeMap(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Multiply elements map", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := map[string]string{
			"Arseniy": "Rust",
			"Roman":  "Golang",
			"Maria":  "Kotlin",
			"Artem":  "Python",
			"Igor":  "OCaml",
		}
		err := internal.SerializeMap(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeMap(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

	t.Run("Empty map", func(t *testing.T) {
        buf := bytes.NewBuffer(make([]byte, 0, 50))
		input := map[string]string{}
		err := internal.SerializeMap(input, buf)
		require.NoError(t, err)

		actual, err := internal.DeserializeMap(buf)
		require.NoError(t, err)

		require.Equal(t, input, actual)
	})

}
