package coll_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestStrings(t *testing.T) {
	testcases := []struct {
		coll.Strings
		i []string
	}{
		{
			Strings: new(coll.Strings).
				Append("hello", "world"),
			i: []string{"hello", "world"},
		},
		{
			Strings: new(coll.Strings).
				Append("hello").
				Append("world"),
			i: []string{"hello", "world"},
		},
	}

	for _, tt := range testcases {
		require.EqualValues(t, tt.i, tt.Strings)
	}
}
