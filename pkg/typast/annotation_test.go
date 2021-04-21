package typast_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestAnnotation_Process(t *testing.T) {
	testcases := []struct {
		TestName    string
		Context     *typgo.Context
		Directives  typast.Directives
		Annotation  *typast.Annotation
		ExpectedErr string
	}{
		{
			Annotation:  &typast.Annotation{},
			ExpectedErr: "mising annotation processor",
		},
		{
			Directives: typast.Directives{{}, {}},
			Annotation: &typast.Annotation{
				Processor: typast.NewProcessor(func(c *typgo.Context, d typast.Directives) error {
					return fmt.Errorf("some-error: %d", len(d))
				}),
			},
			ExpectedErr: "some-error: 2",
		},
		{
			Directives: typast.Directives{{}, {}},
			Annotation: &typast.Annotation{
				Filter: typast.NewFilter(func(d *typast.Directive) bool {
					return false
				}),
				Processor: typast.NewProcessor(func(c *typgo.Context, d typast.Directives) error {
					return fmt.Errorf("some-error: %d", len(d))
				}),
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

func TestAnnotation_ImplementAnnotator(t *testing.T) {
	a := &typast.Annotation{}
	require.Equal(t, a, a.Annotate())
}
