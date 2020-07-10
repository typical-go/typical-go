package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	someFunc   = &typast.Decl{Name: "someFunc", Pkg: "somePkg", Type: typast.Function}
	someFunc2  = &typast.Decl{Name: "someFunc2", Pkg: "somePkg", Type: typast.Function}
	someFunc3  = &typast.Decl{Name: "someFunc3", Pkg: "somePkg", Type: typast.Function}
	someFunc4  = &typast.Decl{Name: "someFunc4", Pkg: "somePkg", Type: typast.Function}
	someStruct = &typast.Decl{Name: "someStruct", Pkg: "somePkg", Type: typast.Struct}
)

func TestCreateCtor(t *testing.T) {
	testcases := []struct {
		testName string
		*typast.Annot
		expected    *typgo.Ctor
		expectedErr string
	}{
		{
			Annot:    &typast.Annot{Decl: someFunc, TagName: "constructor"},
			expected: nil,
		},
		{
			Annot: &typast.Annot{Decl: someFunc, TagName: "ctor"},
			expected: &typgo.Ctor{
				Annot: &typast.Annot{Decl: someFunc, TagName: "ctor"},
			},
		},
		{
			Annot: &typast.Annot{Decl: someFunc, TagName: "ctor", TagAttrs: []byte(`{"name": "noname"}`)},
			expected: &typgo.Ctor{
				Annot: &typast.Annot{Decl: someFunc, TagName: "ctor", TagAttrs: []byte(`{"name": "noname"}`)},
				Param: typgo.CtorParam{
					Name: "noname",
				},
			},
		},
		{
			Annot:       &typast.Annot{Decl: someFunc, TagName: "ctor", TagAttrs: []byte(`{invalid-json`)},
			expectedErr: "someFunc: invalid character 'i' looking for beginning of object key string",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			ctor, err := typgo.ParseCtor(tt.Annot)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, ctor)
			}
		})
	}
}
