package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestContext_CreateImports(t *testing.T) {
	testcases := []struct {
		TestName string
		*typast.Context
		ProjectPkg string
		More       []string
		Expected   []string
	}{
		{
			Context: &typast.Context{
				Dirs: []string{"dir1", "dir2"},
			},
			ProjectPkg: "myproject",
			More:       []string{"github.com/x/x"},
			Expected: []string{
				"myproject/dir1",
				"myproject/dir2",
				"github.com/x/x",
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t,
				tt.Expected,
				tt.CreateImports(tt.ProjectPkg, tt.More...),
			)
		})
	}
}
