package typgen_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestAnnotation_Process(t *testing.T) {
	testcases := []struct {
		TestName    string
		Context     *typgo.Context
		Directives  []*typgen.Directive
		Annotation  *typgen.Annotation
		ExpectedErr string
	}{
		{
			Annotation:  &typgen.Annotation{},
			ExpectedErr: "mising annotation processor",
		},
		{
			Directives: []*typgen.Directive{{}, {}},
			Annotation: &typgen.Annotation{
				ProcessFn: func(c *typgo.Context, d []*typgen.Directive) error {
					return fmt.Errorf("some-error: %d", len(d))
				},
			},
			ExpectedErr: "some-error: 2",
		},
		{
			Directives: []*typgen.Directive{{}, {}},
			Annotation: &typgen.Annotation{
				Filter: typgen.NewFilter(func(d *typgen.Directive) bool {
					return false
				}),
				ProcessFn: func(c *typgo.Context, d []*typgen.Directive) error {
					return fmt.Errorf("some-error: %d", len(d))
				},
			},
			ExpectedErr: "some-error: 0",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			err := tt.Annotation.Process(tt.Context, tt.Directives)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
