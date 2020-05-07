package typannot_test

import "github.com/typical-go/typical-go/pkg/typast"

var (
	someFunc  = &typast.Decl{Name: "someFunc", Pkg: "somePkg", Type: typast.Function}
	someFunc2 = &typast.Decl{Name: "someFunc2", Pkg: "somePkg", Type: typast.Function}
	someFunc3 = &typast.Decl{Name: "someFunc3", Pkg: "somePkg", Type: typast.Function}
	someFunc4 = &typast.Decl{Name: "someFunc4", Pkg: "somePkg", Type: typast.Function}

	someStruct = &typast.Decl{Name: "someStruct", Pkg: "somePkg", Type: typast.Struct}
)
