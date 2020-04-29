package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestCreateAnnotation(t *testing.T) {
	testcases := []struct {
		testName      string
		decl          *typast.Decl
		raw           string
		expected      *typast.Annotation
		expectedError string
	}{
		{
			testName: "tag only",
			raw:      `@autowire`,
			expected: &typast.Annotation{
				TagName: "autowire",
			},
		},
		{
			testName: "tag only with space",
			raw:      `@  autowire  `,
			expected: &typast.Annotation{
				TagName: "autowire",
			},
		},
		{
			testName: "with attribute",
			raw:      `@mock{"pkg":"mock2"}`,
			expected: &typast.Annotation{
				TagName:  "mock",
				TagAttrs: `{"pkg":"mock2"}`,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			annotation, err := typast.CreateAnnotation(tt.decl, tt.raw)
			if tt.expectedError != "" {
				require.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, annotation)
		})
	}
}
