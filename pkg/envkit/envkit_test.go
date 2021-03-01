package envkit_test

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/envkit"
)

func TestSetenv(t *testing.T) {
	m := envkit.Map{
		"key1": "value1",
		"key2": "value2",
	}
	defer envkit.Unsetenv(m)
	envkit.Setenv(m)
	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))
}

func TestNilSafe(t *testing.T) {
	envkit.Setenv(nil)
	envkit.Unsetenv(nil)

	var m envkit.Map = nil
	require.Equal(t, []string(nil), m.SortedKeys())
}

func TestWrite(t *testing.T) {
	m := envkit.Map{
		"key1": "value1",
		"key2": "value2",
	}
	t.Run("success", func(t *testing.T) {
		var b strings.Builder
		require.NoError(t, envkit.Save(m, &b))
		require.Equal(t, "key1=value1\nkey2=value2\n", b.String())
	})
	t.Run("bad writer", func(t *testing.T) {
		require.EqualError(t, envkit.Save(m, &badWriter{}), "bad-writer")
	})
}

func TestSaveFile(t *testing.T) {
	require.NoError(t, envkit.SaveFile(envkit.Map{
		"key1": "value1",
		"key2": "value2",
	}, "some-target"))
	defer os.Remove("some-target")

	b, _ := ioutil.ReadFile("some-target")
	require.Equal(t, "key1=value1\nkey2=value2\n", string(b))
}

func TestReadFile(t *testing.T) {
	ioutil.WriteFile("some-dotenv", []byte("key1=value1\nkey2=value2\n"), 0777)
	defer os.Remove("some-dotenv")
	m, err := envkit.ReadFile("some-dotenv")
	require.NoError(t, err)
	require.Equal(t, envkit.Map{
		"key1": "value1",
		"key2": "value2",
	}, m)
}

func TestReadFile_Error(t *testing.T) {
	_, err := envkit.ReadFile("not-exist")
	require.EqualError(t, err, "open not-exist: no such file or directory")
}

func TestEnvconfig(t *testing.T) {
	type Specification struct {
		RequiredVar string `required:"true"`
	}

	envkit.Setenv(envkit.Map{
		"MYAPP_REQUIREDVAR": "",
	})
	var s Specification
	require.EqualError(t, envconfig.Process("myapp", &s), "required key MYAPP_REQUIREDVAR missing value")
}

type badWriter struct{}

func (*badWriter) Write(p []byte) (n int, err error) { return -1, errors.New("bad-writer") }
