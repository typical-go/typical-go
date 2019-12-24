package common_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestInterfaces_Append(t *testing.T) {
	testcases := []struct {
		*common.Interfaces
		slice []interface{}
	}{
		{
			Interfaces: new(common.Interfaces).
				Append("some-item", 88, 3.14),
			slice: []interface{}{"some-item", 88, 3.14},
		},
		{
			Interfaces: new(common.Interfaces).
				Append("some-item").
				Append(88).
				Append(3.14),
			slice: []interface{}{"some-item", 88, 3.14},
		},
	}
	for _, tt := range testcases {
		require.EqualValues(t, tt.slice, tt.Interfaces.Slice())
	}
}
