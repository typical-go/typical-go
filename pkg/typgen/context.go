package typgen

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Context struct {
		*typgo.Context
		InitAliasGen *AliasGenerator
		InitFuncBody CodeLines
		MappedCoders map[*File][]Coder
		Annotations  []*Annotation
	}
)

var (
	DefaultInitImports = map[string]string{
		"github.com/typical-go/typical-go/pkg/typapp": "",
	}
	InfoSignature = CodeLine("// DO NOT EDIT. Code-generated file\n")
)

func NewContext(c *typgo.Context, annots []*Annotation) *Context {
	return &Context{
		Context:      c,
		InitAliasGen: NewAliasGenerator(DefaultInitImports),
		Annotations:  annots,
		MappedCoders: make(map[*File][]Coder),
	}
}

func (c *Context) AppendInit(s string) {
	c.InitFuncBody = append(c.InitFuncBody, s)
}

func (i *Context) AppendInitf(format string, args ...interface{}) {
	i.AppendInit(fmt.Sprintf(format, args...))
}

func (i *Context) ProvideConstructor(name, importPath, constructor string) {
	alias := i.InitAliasGen.Generate(importPath)
	s := fmt.Sprintf(`typapp.Provide("%s", %s.%s)`, name, alias, constructor)
	i.AppendInit(s)
}

func (i *Context) AppendFileCoder(file *File, coder Coder) {
	i.MappedCoders[file] = append(i.MappedCoders[file], coder)
}

func (i *Context) WriteInitFile(target string) error {
	coder := Coders{
		&File{
			Path:    target,
			Name:    PackageName(target),
			Imports: i.InitAliasGen.Imports(),
		},
		InfoSignature,
		&Function{
			Name: "init",
			Body: i.InitFuncBody,
		},
	}
	return WriteCoder(i.Context, coder, target)
}

func (i *Context) WriteFile(f *File, target string) error {
	coder := Coders{
		&File{
			Path:    target,
			Name:    PackageName(target),
			Imports: f.Imports,
		},
		CodeLine(InfoSignature),
	}
	coder = append(coder, i.MappedCoders[f]...)
	return WriteCoder(i.Context, coder, target)
}
