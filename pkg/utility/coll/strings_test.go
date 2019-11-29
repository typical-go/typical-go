package coll_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestStrings_Append(t *testing.T) {
	testcases := []struct {
		*coll.Strings
		slice  []string
		sorted *coll.Strings
	}{
		{
			Strings: new(coll.Strings).Append("hello", "world"),
			slice:   []string{"hello", "world"},
			sorted:  coll.NewStrings("hello", "world"),
		},
		{
			Strings: new(coll.Strings).Append("hello").Append("world"),
			slice:   []string{"hello", "world"},
			sorted:  coll.NewStrings("hello", "world"),
		},
		{
			Strings: coll.NewStrings("aaa", "ccc", "bbb"),
			slice:   []string{"aaa", "ccc", "bbb"},
			sorted:  coll.NewStrings("aaa", "bbb", "ccc"),
		},
	}
	for _, tt := range testcases {
		require.EqualValues(t, tt.slice, tt.Strings.Slice())
		require.EqualValues(t, tt.sorted, tt.Sort())
	}
}
