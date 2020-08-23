package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestIsPublic(t *testing.T) {
	testnames := []struct {
		TestName string
		Type     typast.Type
		Expected bool
	}{
		{
			Type:     &typast.FuncDecl{Name: "someFunc"},
			Expected: false,
		},
		{
			Type:     &typast.FuncDecl{Name: "SomeFunc"},
			Expected: true,
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typast.IsPublic(tt.Type))
		})
	}
}
