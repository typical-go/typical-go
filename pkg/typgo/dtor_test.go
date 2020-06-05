package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCreateDtor(t *testing.T) {
	testcases := []struct {
		testName string
		*typast.Annot
		expected *typgo.Dtor
	}{
		{
			Annot: &typast.Annot{Decl: someFunc, TagName: "dtor"},
			expected: &typgo.Dtor{
				Annot: &typast.Annot{Decl: someFunc, TagName: "dtor"},
			},
		},
		{
			Annot:    &typast.Annot{Decl: someFunc, TagName: "wrong-tag"},
			expected: nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, typgo.CreateDtor(tt.Annot))
		})
	}
}
