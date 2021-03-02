package typgo_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestBuildTool_Environment(t *testing.T) {
	var out strings.Builder
	typgo.BuildTool(&typgo.Descriptor{
		Environment: typgo.Environment{
			"key1": "value1",
			"key2": "value2",
		},
		Stdout: &out,
	})

	defer os.Clearenv()

	require.Equal(t, "> load environment\nkey1, key2\n", out.String())
	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))
}

func TestBuildTool_EnvironmentError(t *testing.T) {
	var out strings.Builder
	typgo.BuildTool(&typgo.Descriptor{
		Environment: typgo.DotEnv("not-found"),
		Stdout:      &out,
	})

	defer os.Clearenv()

	require.Equal(t, "> load environment: open not-found: no such file or directory\n", out.String())

}
