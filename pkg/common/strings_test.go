package common_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestStrings_Append(t *testing.T) {
	testcases := []struct {
		*common.Strings
		slice  []string
		sorted *common.Strings
	}{
		{
			Strings: new(common.Strings).Append("hello", "world"),
			slice:   []string{"hello", "world"},
			sorted:  common.NewStrings("hello", "world"),
		},
		{
			Strings: new(common.Strings).Append("hello").Append("world"),
			slice:   []string{"hello", "world"},
			sorted:  common.NewStrings("hello", "world"),
		},
		{
			Strings: common.NewStrings("aaa", "ccc", "bbb"),
			slice:   []string{"aaa", "ccc", "bbb"},
			sorted:  common.NewStrings("aaa", "bbb", "ccc"),
		},
	}
	for _, tt := range testcases {
		require.EqualValues(t, tt.slice, tt.Strings.Slice())
		require.EqualValues(t, tt.sorted, tt.Sort())
	}
}

func TestStrings_Reverse(t *testing.T) {
	testcases := []struct {
		*common.Strings
		reversed []string
	}{
		{common.NewStrings("a", "b", "c"), []string{"c", "b", "a"}},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.reversed, tt.Reverse().Slice())
	}
}
