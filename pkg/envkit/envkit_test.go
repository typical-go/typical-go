package envkit_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/envkit"
)

func TestSortedKey(t *testing.T) {
	testcases := []struct {
		TestName string
		Raw      string
		Expected map[string]string
	}{
		{
			Raw: "key1=value1",
			Expected: map[string]string{
				"key1": "value1",
			},
		},
		{
			Raw: "key1=value1\n\nkey2=value2\nkey3=value3\n\n",
			Expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			r := strings.NewReader(tt.Raw)
			require.Equal(t, tt.Expected, envkit.Read(r))
		})
	}
}

func TestSetenv(t *testing.T) {
	m := map[string]string{
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

	var m map[string]string
	require.Equal(t, []string(nil), envkit.SortedKeys(m))
}

func TestWrite(t *testing.T) {
	m := map[string]string{
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
	require.NoError(t, envkit.SaveFile(map[string]string{
		"key1": "value1",
		"key2": "value2",
	}, "some-target"))
	defer os.Remove("some-target")

	b, _ := os.ReadFile("some-target")
	require.Equal(t, "key1=value1\nkey2=value2\n", string(b))
}

func TestReadFile(t *testing.T) {
	os.WriteFile("some-dotenv", []byte("key1=value1\nkey2=value2\n"), 0777)
	defer os.Remove("some-dotenv")
	m, err := envkit.ReadFile("some-dotenv")
	require.NoError(t, err)
	require.Equal(t, map[string]string{
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

	envkit.Setenv(map[string]string{
		"MYAPP_REQUIREDVAR": "",
	})
	var s Specification
	require.EqualError(t, envconfig.Process("myapp", &s), "required key MYAPP_REQUIREDVAR missing value")
}

type badWriter struct{}

func (*badWriter) Write(p []byte) (n int, err error) { return -1, errors.New("bad-writer") }
