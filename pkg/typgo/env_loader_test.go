package typgo_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDotEnv(t *testing.T) {
	os.WriteFile(".env", []byte("key1=value1\nkey2=value2\n"), 0777)
	defer os.Remove(".env")

	m, err := typgo.DotEnv(".env").EnvLoad()
	require.NoError(t, err)

	require.Equal(t, map[string]string{
		"key1": "value1",
		"key2": "value2",
	}, m)
}

func TestEnvironment(t *testing.T) {
	m, err := typgo.Environment{
		"key1": "value1",
		"key2": "value2",
	}.EnvLoad()
	require.NoError(t, err)
	require.Equal(t, map[string]string{
		"key1": "value1",
		"key2": "value2",
	}, m)
}
