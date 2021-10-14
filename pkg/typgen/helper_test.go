package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestCreateTargetDir(t *testing.T) {
	testCases := []struct {
		TestName   string
		Annotation *typgen.Annotation
		Suffix     string
		Expected   string
	}{
		{
			Annotation: &typgen.Annotation{
				Decl: &typgen.Decl{
					File: &typgen.File{
						Path: ".",
					},
				},
			},
			Expected: "internal/generated",
		},
		{
			Annotation: &typgen.Annotation{
				Decl: &typgen.Decl{
					File: &typgen.File{
						Path: "internal/app/service/file.go",
					},
				},
			},
			Suffix:   "mock",
			Expected: "internal/generated/app/service_mock",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			dir := typgen.CreateTargetDir(tt.Annotation, tt.Suffix)
			require.Equal(t, tt.Expected, dir)
		})
	}
}
