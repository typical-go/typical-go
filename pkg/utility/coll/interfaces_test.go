package coll_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/utility/coll"
)

func TestInterfaces(t *testing.T) {
	testcases := []struct {
		*coll.Interfaces
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
		require.EqualValues(t, tt.i, *tt.Interfaces)
	}
}

func TestInterfaces2(t *testing.T) {
	t.Run("Non-pointer definition", func(t *testing.T) {
		var s coll.Interfaces
		s.Append("hello").Append(123456)
		require.EqualValues(t, []interface{}{"hello", 123456}, s)
	})
}
