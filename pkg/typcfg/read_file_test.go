package typcfg_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

func TestReadFile(t *testing.T) {
	testcases := []struct {
		raw            string
		expectedResult map[string]string
		expectedErr    string
	}{
		{
			raw: "key1=value1\nkey2=value2",
			expectedResult: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}

	for i, tt := range testcases {
		src := fmt.Sprintf("test-%d.env", i)

		ioutil.WriteFile(src, []byte(tt.raw), 0777)
		defer os.Remove(src)

		res, err := typcfg.ReadFile(src)
		if tt.expectedErr != "" {
			require.EqualError(t, err, tt.expectedErr)
		} else {
			require.NoError(t, err)
		}

		require.Equal(t, tt.expectedResult, res)
	}

}
