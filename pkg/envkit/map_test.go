package envkit_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/envkit"
)

func TestMap(t *testing.T) {
	testcases := []struct {
		TestName string
		Raw      string
		Expected envkit.Map
	}{
		{
			Raw: "key1=value1",
			Expected: envkit.Map{
				"key1": "value1",
			},
		},
		{
			Raw: "key1=value1\n\nkey2=value2\nkey3=value3\n\n",
			Expected: envkit.Map{
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
