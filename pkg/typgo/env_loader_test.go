package typgo_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/oskit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDotEnv(t *testing.T) {
	ioutil.WriteFile(".env", []byte("key1=value1\nkey2=value2\n"), 0777)
	defer os.Remove(".env")

	var out strings.Builder
	defer oskit.PatchStdout(&out)()

	require.NoError(t, typgo.DotEnv(".env").EnvLoad())
	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))
	require.Equal(t, "Load environment from '.env': [key1 key2]\n\n", out.String())
}

func TestEnvMap(t *testing.T) {
	var out strings.Builder
	defer oskit.PatchStdout(&out)()

	require.NoError(t, typgo.EnvMap{
		"key1": "value1",
		"key2": "value2",
	}.EnvLoad())
	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))
	require.Equal(t, "Load environment: [key1 key2]\n\n", out.String())
}
