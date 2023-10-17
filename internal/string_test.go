package internal

//// Check Serialize - Deserialize
//func TestSerializeDeserializeString(t *testing.T) {
//	t.Run("hello", func(t *testing.T) {
//		input := "hello"
//		bytes, err := SerializeString(input)
//		require.NoError(t, err)
//
//		actual, err := DeserializeString(bytes)
//		require.NoError(t, err)
//
//		require.Equal(t, input, actual)
//	})
//
//	t.Run("World Of Tanks", func(t *testing.T) {
//		input := "World Of Tanks"
//		bytes, err := SerializeString(input)
//		require.NoError(t, err)
//
//		actual, err := DeserializeString(bytes)
//		require.NoError(t, err)
//
//		require.Equal(t, input, actual)
//	})
//
//	t.Run("LOL", func(t *testing.T) {
//		input := "LOL"
//		bytes, err := SerializeString(input)
//		require.NoError(t, err)
//
//		actual, err := DeserializeString(bytes)
//		require.NoError(t, err)
//
//		require.Equal(t, input, actual)
//	})
//}
