package typcfg_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestReadConfig(t *testing.T) {
	testcases := []struct {
		testName string
		raw      string
		expected map[string]string
	}{
		{
			raw: "key1=value1\nkey2=value2",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			raw: "key1=value1\n\n\nkey2=value2",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t,
				tt.expected,
				typcfg.ReadConfig(strings.NewReader(tt.raw)),
			)
		})
	}
}

func TestWriteConfig(t *testing.T) {
	testcases := []struct {
		testName string
		typcfg.Configurer
		before      string
		expected    string
		expectedErr string
	}{
		{
			Configurer: &typcfg.Configuration{
				Name: "TEST",
				Spec: &someSpec{},
			},
			expected: "TEST_FIELD1=defaulValue1\nTEST_FIELD2=defaulValue2\n",
		},
		{
			Configurer: &typcfg.Configuration{
				Name: "TEST",
				Spec: &someSpec{},
			},
			before:   "XXXX=XXXX",
			expected: "XXXX=XXXX\nTEST_FIELD1=defaulValue1\nTEST_FIELD2=defaulValue2\n",
		},
	}

	for i, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			dest := fmt.Sprintf("write%d.env", i)
			defer os.Remove(dest)

			if tt.before != "" {
				ioutil.WriteFile(dest, []byte(tt.before), 0777)
			}

			err := typcfg.WriteConfig(dest, tt.Configurer)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}

			b, _ := ioutil.ReadFile(dest)
			require.Equal(t, tt.expected, string(b))
		})
	}
}

type someSpec struct {
	Field1 string `default:"defaulValue1"`
	Field2 string `default:"defaulValue2"`
}
