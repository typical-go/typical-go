package typgo_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/typical-go/typical-go/pkg/oskit"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestCli_Before(t *testing.T) {
	app := typgo.Cli(&typgo.BuildSys{})
	require.NoError(t, app.Before(&cli.Context{}))
}

func TestCli_Before_Setenv(t *testing.T) {
	ioutil.WriteFile(".env", []byte("key1=value1\nkey2=value2\n"), 0777)
	defer os.Remove(".env")

	os.Clearenv()
	defer os.Clearenv()

	var out strings.Builder
	defer oskit.PatchStdout(&out)()

	app := typgo.Cli(&typgo.BuildSys{})
	require.NoError(t, app.Before(&cli.Context{}))

	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))

	require.Equal(t, "Load environment '.env' [key1 key2]\n", out.String())
}
