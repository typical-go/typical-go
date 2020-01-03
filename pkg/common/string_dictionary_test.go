package common_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestStringDictionary(t *testing.T) {
	testcases := []struct {
		*common.StringDictionary
		slice []*common.StringKV
	}{
		{
			StringDictionary: new(common.StringDictionary).
				Add("key-1", "value-1").
				Add("key-2", "value-2").
				Add("key-3", "value-3"),
			slice: []*common.StringKV{
				&common.StringKV{"key-1", "value-1"},
				&common.StringKV{"key-2", "value-2"},
				&common.StringKV{"key-3", "value-3"},
			},
		},
	}
	for i, tt := range testcases {
		require.EqualValues(t, tt.slice, tt.StringDictionary.Slice(), i)
	}
}
