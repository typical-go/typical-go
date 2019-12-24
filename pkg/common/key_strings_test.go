package common_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestKeyStrings(t *testing.T) {
	testcases := []struct {
		*common.KeyStrings
		slice []*common.KeyString
	}{
		{
			KeyStrings: new(common.KeyStrings).Append(
				&common.KeyString{"key-1", "string-1"},
				&common.KeyString{"key-2", "string-2"},
				&common.KeyString{"key-3", "string-3"},
			),
			slice: []*common.KeyString{
				&common.KeyString{"key-1", "string-1"},
				&common.KeyString{"key-2", "string-2"},
				&common.KeyString{"key-3", "string-3"},
			},
		},
		{
			KeyStrings: new(common.KeyStrings).
				Append(&common.KeyString{"key-1", "string-1"}).
				Append(&common.KeyString{"key-2", "string-2"}).
				Append(&common.KeyString{"key-3", "string-3"}),
			slice: []*common.KeyString{
				&common.KeyString{"key-1", "string-1"},
				&common.KeyString{"key-2", "string-2"},
				&common.KeyString{"key-3", "string-3"},
			},
		},
	}
	for i, tt := range testcases {
		require.EqualValues(t, tt.slice, tt.KeyStrings.Slice(), i)
	}
}
