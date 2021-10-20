package typgen

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Context struct {
		*typgo.Context
		InitAlias    *AliasGenerator
		InitFuncBody []string
		Annotations  []*Annotation
	}
)

var (
	defaultInitFileImports = map[string]string{
		"github.com/typical-go/typical-go/pkg/typapp": "",
	}
)

func NewContext(c *typgo.Context, annots []*Annotation) *Context {
	return &Context{
		Context:     c,
		InitAlias:   NewAliasGenerator(defaultInitFileImports),
		Annotations: annots,
	}
}

func (i *Context) WriteInitFile(c *typgo.Context, target string) error {
	coder := i.createInitCoder(target)
	return WriteCoder(c, coder, target)
}

func (i *Context) createInitCoder(target string) Coder {
	return Coders{
		&File{
			Name:   PackageName(target),
			Import: i.InitAlias.Imports(),
		},
		Comment("DO NOT EDIT. Code-generated file."),
		&Function{
			Name: "init",
			Body: i.InitFuncBody,
		},
	}
}

func (i *Context) AppendInit(s string) {
	i.InitFuncBody = append(i.InitFuncBody, s)
}

func (i *Context) AppendInitf(format string, args ...interface{}) {
	i.AppendInit(fmt.Sprintf(format, args...))
}

func (i *Context) ProvideConstructor(name, importPath, constructor string) {
	alias := i.InitAlias.Generate(importPath)
	s := fmt.Sprintf(`typapp.Provide("%s", %s.%s)`, name, alias, constructor)
	i.AppendInit(s)
}
