package coll_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestInterfaces(t *testing.T) {
	testcases := []struct {
		coll.Interfaces
		i []interface{}
	}{
		{
			Interfaces: new(coll.Interfaces).
				Append("some-item", 88, 3.14),
			i: []interface{}{"some-item", 88, 3.14},
		},
		{
			Interfaces: new(coll.Interfaces).
				Append("some-item").
				Append(88).
				Append(3.14),
			i: []interface{}{"some-item", 88, 3.14},
		},
	}

	for _, tt := range testcases {
		require.EqualValues(t, tt.i, tt.Interfaces)
	}
}
