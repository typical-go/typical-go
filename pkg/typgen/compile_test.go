package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestParseRawAnnot(t *testing.T) {
	testcases := []struct {
		TestName         string
		Raw              string
		ExpectedTagName  string
		ExpectedTagAttrs string
	}{
		{
			TestName:        "tag only",
			Raw:             `@tag1`,
			ExpectedTagName: "@tag1",
		},
		{
			TestName:        "tag only with space",
			Raw:             `@tag2 extra1`,
			ExpectedTagName: "@tag2",
		},
		{
			TestName:         "with attribute",
			Raw:              `@tag3("name":"wire1")`,
			ExpectedTagName:  "@tag3",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
		{
			TestName:         "there is space between tagname and attribute",
			Raw:              `@tag4 ("name":"wire1")`,
			ExpectedTagName:  "@tag4",
			ExpectedTagAttrs: `"name":"wire1"`,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			tagName, tagAttrs := typgen.ParseRawAnnot(tt.Raw)
			require.Equal(t, tt.ExpectedTagName, tagName)
			require.Equal(t, tt.ExpectedTagAttrs, tagAttrs)
		})
	}
}
