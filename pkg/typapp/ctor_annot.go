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
		Target   string        // By default is `internal/generated/ctor/ctor.go`
		Filter   typgen.Filter // By default is annotated by `@ctor` and has public acces
		imports  *tmplkit.Imports
		initFunc []fmt.Stringer
	}
)

var (
	DefaultCtorTag    = "@ctor"
	DefaultCtorFilter = typgen.Filters{
		&typgen.TagNameFilter{DefaultCtorTag},
		&typgen.PublicFilter{},
	}
	DefaultCtorTarget = "internal/generated/ctor/ctor.go"
)

//
// CtorAnnot
//

var _ typgen.Processor = (*CtorAnnot)(nil)

func (a *CtorAnnot) Process(c *typgo.Context, directives []*typgen.Directive) error {
	return a.Annotation().Process(c, directives)
}

func (a *CtorAnnot) Annotation() *typgen.Annotation {
	if a.Filter == nil {
		a.Filter = DefaultCtorFilter
	}
	if a.Target == "" {
		a.Target = DefaultCtorTarget
	}

	return &typgen.Annotation{
		Filter:    a.Filter,
		ProcessFn: a.process,
	}
}

func (a *CtorAnnot) appendImport(pkg string) string {
	return a.Imports().AppendWithAlias(pkg)
}

func (a *CtorAnnot) process(c *typgo.Context, directives []*typgen.Directive) error {
	for _, d := range directives {
		if err := a.GenerateCode(c, d); err != nil {
			return err
		}
	}

	dest := filepath.Dir(a.Target)
	sourceCode := tmplkit.SourceCode{
		Signature: tmplkit.Signature{},
		Package:   filepath.Base(dest),
		Imports:   a.Imports(),
		LineCodes: []fmt.Stringer{
			tmplkit.InitFunction(a.initFunc),
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

func (a *CtorAnnot) GenerateCode(c *typgo.Context, d *typgen.Directive) error {

	var lines []string

	switch d.Type.(type) {
	case *typgen.FuncDecl:
		lines = append(lines, a.generateCtorForFunc(d))
	case *typgen.StructDecl:
		lines = append(lines, "// TODO")
	default:
		lines = append(lines, fmt.Sprintf("// '%s' is not supported", d.GetName()))
	}

	for _, line := range lines {
		a.initFunc = append(a.initFunc, tmplkit.LineCode(line))
	}
	return nil
}

func (a *CtorAnnot) generateCtorForFunc(d *typgen.Directive) string {
	currPackagePath := fmt.Sprintf("%s/%s", typgo.ProjectPkg, filepath.Dir(d.File.Path))
	alias := a.appendImport(currPackagePath)

	funcDecl := d.Type.(*typgen.FuncDecl)

	name := d.TagParam.Get("name")
	if !funcDecl.IsMethod() {
		return fmt.Sprintf(`typapp.Provide("%s", %s.%s)`, name, alias, d.GetName())
	}
	return fmt.Sprintf("// Method '%s' is not supported", d.GetName())
}

func (a *CtorAnnot) InitFunc() []fmt.Stringer {
	return a.initFunc
}

func (a *CtorAnnot) Imports() *tmplkit.Imports {
	if a.imports == nil {
		a.imports = tmplkit.NewImports(nil)
		a.imports.Map["github.com/typical-go/typical-go/pkg/typapp"] = ""
	}
	return a.imports
}
