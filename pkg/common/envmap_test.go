package common_test

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestEnvMap(t *testing.T) {
	testcases := []struct {
		TestName string
		Raw      string
		Expected common.EnvMap
	}{
		{
			Raw: "key1=value1",
			Expected: common.EnvMap{
				"key1": "value1",
			},
		},
		{
			Raw: "key1=value1\n\nkey2=value2\nkey3=value3\n\n",
			Expected: common.EnvMap{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			r := strings.NewReader(tt.Raw)
			require.Equal(t, tt.Expected, common.CreateEnvMap(r))
		})
	}
}

func TestLoadEnv(t *testing.T) {
	filename := "env1"
	ioutil.WriteFile(filename, []byte("key1=value1"), 0777)
	defer os.Remove(filename)
	defer os.Clearenv()

	require.NoError(t, common.LoadEnv(filename))
	require.Equal(t, "value1", os.Getenv("key1"))
}

func TestLoadEnv_NoFileExit(t *testing.T) {
	require.NoError(t, common.LoadEnv("not-exist"))
}

func TestEnvMap_Setenv(t *testing.T) {
	m := common.EnvMap{
		"key1": "value1",
		"key2": "value2",
	}
	defer common.Unsetenv(m)
	common.Setenv(m)
	require.Equal(t, "value1", os.Getenv("key1"))
	require.Equal(t, "value2", os.Getenv("key2"))
}

func TestEnvMap_Write(t *testing.T) {
	m := common.EnvMap{
		"key1": "value1",
		"key2": "value2",
	}
	t.Run("success", func(t *testing.T) {
		var b strings.Builder
		require.NoError(t, m.Save(&b))
		require.Equal(t, "key1=value1\nkey2=value2\n", b.String())
	})
	t.Run("bad writer", func(t *testing.T) {
		require.EqualError(t, m.Save(&badWriter{}), "bad-writer")
	})
}

type badWriter struct{}

func (*badWriter) Write(p []byte) (n int, err error) { return -1, errors.New("bad-writer") }
