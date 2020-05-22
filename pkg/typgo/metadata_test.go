package typgo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestOpenMetadata_NotExist(t *testing.T) {
	db, err := typgo.OpenMetadata("not-exist")
	defer os.Remove("not-exist")

	require.NoError(t, err)
	require.Equal(t, &typgo.Metadata{
		Path:  "not-exist",
		Extra: map[string]interface{}{},
	}, db)

	b, _ := ioutil.ReadFile("not-exist")
	require.Equal(t, `{}`, string(b))
}

func TestOpenMetadata(t *testing.T) {
	testcases := []struct {
		testName    string
		path        string
		data        string
		expected    *typgo.Metadata
		expectedErr string
	}{
		{
			path: "test.json",
			data: `{"path": "test.json", "extra": {"key-1":"value-1", "key-2":"value-2"}}`,
			expected: &typgo.Metadata{
				Path: "test.json",
				Extra: map[string]interface{}{
					"key-1": "value-1",
					"key-2": "value-2",
				},
			},
		},
		{
			testName:    "broken json",
			path:        "test.json",
			data:        `{invalid-json`,
			expectedErr: "invalid character 'i' looking for beginning of object key string",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			ioutil.WriteFile(tt.path, []byte(tt.data), 0777)
			defer os.Remove(tt.path)

			db, err := typgo.OpenMetadata(tt.path)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, db)
		})

	}

}
