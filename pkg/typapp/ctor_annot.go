package typapp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/tmplkit"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// CtorAnnot handle @ctor annotation
	// e.g. `@ctor (name:"NAME")`
	CtorAnnot struct {
		TagName string // By default is `@ctor`
		Target  string // By default is `internal/generated/ctor/ctor.go`
	}
)

//
// CtorAnnot
//

var _ typgen.Processor = (*CtorAnnot)(nil)

func (a *CtorAnnot) Process(c *typgo.Context, directives []*typgen.Directive) error {
	return a.Annotation().Process(c, directives)
}

func (a *CtorAnnot) Annotation() *typgen.Annotation {
	if a.TagName == "" {
		a.TagName = "@ctor"
	}
	if a.Target == "" {
		a.Target = "internal/generated/ctor/ctor.go"
	}

	return &typgen.Annotation{
		Filter: typgen.Filters{
			&typgen.TagNameFilter{a.TagName},
			&typgen.PublicFilter{},
			&typgen.FuncFilter{},
		},
		ProcessFn: a.process,
	}
}

func (a *CtorAnnot) process(c *typgo.Context, directives []*typgen.Directive) error {
	imports := tmplkit.NewImports(nil)
	var lineCodes []fmt.Stringer
	for _, directive := range directives {
		alias := imports.AppendWithAlias(directive.Package())
		lineCodes = append(lineCodes, generateCodeForCtor(alias, directive))

	}
	imports.Map["github.com/typical-go/typical-go/pkg/typapp"] = ""

	dest := filepath.Dir(a.Target)
	sourceCode := tmplkit.SourceCode{
		Signature: tmplkit.Signature{TagName: a.TagName},
		Package:   filepath.Base(dest),
		Imports:   imports,
		LineCodes: []fmt.Stringer{
			tmplkit.InitFunction(lineCodes),
		},
	}

	os.MkdirAll(dest, 0777)
	c.Infof("Generate @ctor to %s\n", a.Target)

	err := ioutil.WriteFile(a.Target, []byte(sourceCode.String()), 0644)
	if err != nil {
		return err
	}
	typgo.GoImports(c, a.Target)
	return nil
}

func generateCodeForCtor(importAlias string, d *typgen.Directive) tmplkit.LineCode {
	name := d.TagParam.Get("name")
	s := fmt.Sprintf("typapp.Provide(\"%s\", %s.%s)", name, importAlias, d.GetName())
	return tmplkit.LineCode(s)
}
