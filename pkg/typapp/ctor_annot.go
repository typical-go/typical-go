package typapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// CtorAnnot handle @ctor annotation
	// e.g. `@ctor (name:"NAME")`
	CtorAnnot struct {
		Target         string        // By default is `internal/generated/ctor/ctor.go`
		Filter         typgen.Filter // By default is annotated by `@ctor` and has public acces
		aliasGenerator *typgen.AliasGenerator
		initLines      []string
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
		ProcessFn: a.GenerateCode,
	}
}

func (a *CtorAnnot) generateAlias(pkg string) string {
	return a.AliasGenerator().Generate(pkg)
}

func (a *CtorAnnot) GenerateCode(c *typgo.Context, directives []*typgen.Directive) error {
	for _, d := range directives {
		switch d.Type.(type) {
		case *typgen.Function:
			a.initLines = append(a.initLines, a.generateCodeForFunc(d))
		case *typgen.Struct:
			a.initLines = append(a.initLines, a.generateCodeForStruct(d))
		default:
			a.initLines = append(a.initLines, a.unsupportedType(d))
		}
	}

	dest := filepath.Dir(a.Target)

	os.MkdirAll(dest, 0777)
	c.Infof("Generate @ctor to %s\n", a.Target)

	err := typgen.WriteSourceCode(a.Target,
		&typgen.File{
			Name:   filepath.Base(dest),
			Import: a.AliasGenerator().Imports(),
		},
		typgen.Comment("DO NOT EDIT. Code-generated file."),
		&typgen.Function{
			Name: "init",
			Body: a.initLines,
		},
	)

	typgo.GoImports(c, a.Target)
	return err
}

func (a *CtorAnnot) generateCodeForFunc(d *typgen.Directive) string {
	currPackagePath := fmt.Sprintf("%s/%s", typgo.ProjectPkg, filepath.Dir(d.File.Path))
	alias := a.generateAlias(currPackagePath)

	funcDecl := d.Type.(*typgen.Function)

	name := d.TagParam.Get("name")
	if !funcDecl.IsMethod() {
		return fmt.Sprintf(`typapp.Provide("%s", %s.%s)`, name, alias, d.GetName())
	}
	return fmt.Sprintf("// Method '%s' is not supported", d.GetName())
}

func (a *CtorAnnot) generateCodeForStruct(d *typgen.Directive) string {
	return "// TODO"
}

func (a *CtorAnnot) unsupportedType(d *typgen.Directive) string {
	return fmt.Sprintf("// '%s' is not supported", d.GetName())
}

func (a *CtorAnnot) AliasGenerator() *typgen.AliasGenerator {
	if a.aliasGenerator == nil {
		a.aliasGenerator = typgen.NewAliasGenerator(nil)
		a.aliasGenerator.Map["github.com/typical-go/typical-go/pkg/typapp"] = ""
	}
	return a.aliasGenerator
}
