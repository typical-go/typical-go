package coll_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestKeyStrings(t *testing.T) {
	testcases := []struct {
		*coll.KeyStrings
		i []coll.KeyString
	}{
		{
			KeyStrings: new(coll.KeyStrings).Append(
				coll.KeyString{"key-1", "string-1"},
				coll.KeyString{"key-2", "string-2"},
				coll.KeyString{"key-3", "string-3"},
			),
			i: []coll.KeyString{
				coll.KeyString{"key-1", "string-1"},
				coll.KeyString{"key-2", "string-2"},
				coll.KeyString{"key-3", "string-3"},
			},
		},
		{
			KeyStrings: new(coll.KeyStrings).
				Append(coll.KeyString{"key-1", "string-1"}).
				Append(coll.KeyString{"key-2", "string-2"}).
				Append(coll.KeyString{"key-3", "string-3"}),
			i: []coll.KeyString{
				coll.KeyString{"key-1", "string-1"},
				coll.KeyString{"key-2", "string-2"},
				coll.KeyString{"key-3", "string-3"},
			},
		},
	}
	for i, tt := range testcases {
		require.EqualValues(t, tt.i, *tt.KeyStrings, i)
	}
}
