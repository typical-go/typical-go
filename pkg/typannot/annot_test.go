package typannot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

func TestCreateAnnotation(t *testing.T) {
	testcases := []struct {
		testName      string
		decl          *typannot.Decl
		raw           string
		expected      *typannot.Annot
		expectedError string
	}{
		{
			testName: "tag only",
			raw:      `@autowire`,
			expected: &typannot.Annot{
				TagName: "autowire",
			},
		},
		{
			testName: "tag only with space",
			raw:      `@  autowire  `,
			expected: &typannot.Annot{
				TagName: "autowire",
			},
		},
		{
			testName: "with attribute",
			raw:      `@mock{"pkg":"mock2"}`,
			expected: &typannot.Annot{
				TagName:  "mock",
				TagAttrs: []byte(`{"pkg":"mock2"}`),
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			annotation, err := typannot.CreateAnnot(tt.decl, tt.raw)
			if tt.expectedError != "" {
				require.EqualError(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expected, annotation)
		})
	}
}

func TestUnmarshall(t *testing.T) {
	testcases := []struct {
		testName string
		*typannot.Annot
		expected    map[string]string
		expectedErr string
	}{
		{
			testName: "",
			Annot: &typannot.Annot{
				TagName:  "mock",
				TagAttrs: []byte(`{"key1":"value1"}`),
			},
			expected: map[string]string{
				"key1": "value1",
			},
		},
		{
			testName: "",
			Annot: &typannot.Annot{
				TagName: "mock",
			},
		},
		{
			testName: "",
			Annot: &typannot.Annot{
				TagName:  "mock",
				TagAttrs: []byte(`{"key1":"value1"`),
			},
			expectedErr: "unexpected end of JSON input",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			var m map[string]string
			err := tt.Unmarshal(&m)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, m)
			}
		})
	}
}

func TestAnnot_CheckFunc(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Annot
		TagName  string
		Expected bool
	}{
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "tagname",
			Expected: true,
		},
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "tagname1",
			Expected: false,
		},
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.StructType{}}},
			TagName:  "tagname",
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.CheckFunc(tt.TagName))
		})
	}
}

func TestAnnot_CheckStruct(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Annot
		TagName  string
		Expected bool
	}{
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.StructType{}}},
			TagName:  "tagname",
			Expected: true,
		},
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.StructType{}}},
			TagName:  "tagname1",
			Expected: false,
		},
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "tagname",
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.CheckStruct(tt.TagName))
		})
	}
}

func TestAnnot_CheckInterface(t *testing.T) {
	testcases := []struct {
		TestName string
		*typannot.Annot
		TagName  string
		Expected bool
	}{
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.InterfaceType{}}},
			TagName:  "tagname",
			Expected: true,
		},
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.InterfaceType{}}},
			TagName:  "tagname1",
			Expected: false,
		},
		{
			Annot:    &typannot.Annot{TagName: "tagname", Decl: &typannot.Decl{Type: &typannot.FuncType{}}},
			TagName:  "tagname",
			Expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.CheckInterface(tt.TagName))
		})
	}
}
