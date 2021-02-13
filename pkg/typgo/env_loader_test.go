package typgo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDotEnv(t *testing.T) {
	ioutil.WriteFile(".env", []byte("key1=value1\nkey2=value2\n"), 0777)
	defer os.Remove(".env")

	c, out := typgo.DummyBuildToolContext()
	require.NoError(t, typgo.DotEnv(".env").EnvLoad(c))
	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))
	require.Equal(t, "some-project:dummy> Read from DotEnv '.env': [key1 key2]\n", out.String())
}

func TestEnvMap(t *testing.T) {
	c, out := typgo.DummyBuildToolContext()
	require.NoError(t, typgo.EnvMap{
		"key1": "value1",
		"key2": "value2",
	}.EnvLoad(c))
	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))
	require.Equal(t, "some-project:dummy> Read from EnvMap: [key1 key2]\n", out.String())
}
