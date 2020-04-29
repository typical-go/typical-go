package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestCreateAnnotation(t *testing.T) {
	testcases := []struct {
		testName string
		decl     *typast.Decl
		raw      string
		*typast.Annotation
	}{
		{
			testName: "tag only",
			raw:      `@autowire`,
			Annotation: &typast.Annotation{
				TagName:  "autowire",
				TagAttrs: map[string]string{},
			},
		},
		{
			testName: "tag only with space",
			raw:      `@  autowire  `,
			Annotation: &typast.Annotation{
				TagName:  "autowire",
				TagAttrs: map[string]string{},
			},
		},
		{
			testName: "with attribute (no quote)",
			raw:      `@mock(pkg=mock2)`,
			Annotation: &typast.Annotation{
				TagName: "mock",
				TagAttrs: map[string]string{
					"pkg": "mock2",
				},
			},
		},

		{
			testName: "with attribute, no quote, space between tagname and bracket",
			raw:      `@mock (pkg=mock2)`,
			Annotation: &typast.Annotation{
				TagName: "mock",
				TagAttrs: map[string]string{
					"pkg": "mock2",
				},
			},
		},
		{
			testName: "with attribute (with quote)",
			raw:      `@mock(pkg="mock2")`,
			Annotation: &typast.Annotation{
				TagName: "mock",
				TagAttrs: map[string]string{
					"pkg": "mock2",
				},
			},
		},
		{
			testName: "with invalid attribute (missing close bracket)",
			raw:      `@mock(pkg="mock2"`,
			Annotation: &typast.Annotation{
				TagName:  "mock",
				TagAttrs: map[string]string{},
			},
		},
		{
			testName: "with attribute (no value)",
			raw:      `@mock(pkg)`,
			Annotation: &typast.Annotation{
				TagName: "mock",
				TagAttrs: map[string]string{
					"pkg": "",
				},
			},
		},
		{
			testName: "multiple key attribute (value with quote)",
			raw:      `@noname(key1="value1" key2="value2")`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			testName: "multiple key attribute (value with no quote)",
			raw:      `@noname(key1=value1 key2=value2)`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			testName: "multiple key attribute",
			raw:      `@noname(key1=value1 key2="value2")`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			testName: "multiple key attribute",
			raw:      `@noname(key1=value1 key2 key3=value3)`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "value1",
					"key2": "",
					"key3": "value3",
				},
			},
		},
		{
			testName: "multiple key attribute",
			raw:      `@noname(key1= key2 key3="")`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "",
					"key2": "",
					"key3": "",
				},
			},
		},
		{
			testName: "multiple key attribute",
			raw:      `@noname(key1="" key2 key3=)`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "",
					"key2": "",
					"key3": "",
				},
			},
		},
		{
			testName: "multiple key attribute",
			raw:      `@noname(key1="" key2 key3)`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "",
					"key2": "",
					"key3": "",
				},
			},
		},
		{
			testName: "multiple key attribute",
			raw:      `@noname(key1="" key2 key3 key4=value4)`,
			Annotation: &typast.Annotation{
				TagName: "noname",
				TagAttrs: map[string]string{
					"key1": "",
					"key2": "",
					"key3": "",
					"key4": "value4",
				},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.Annotation, typast.CreateAnnotation(tt.decl, tt.raw))
		})
	}
}
