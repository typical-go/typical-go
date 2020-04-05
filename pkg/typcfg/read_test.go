package typcfg_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestReadFile(t *testing.T) {
	testcases := []struct {
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
		require.Equal(t, tt.expected, typcfg.Read(strings.NewReader(tt.raw)))
	}

}
