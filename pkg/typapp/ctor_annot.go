package typapp

import (
	"fmt"
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
		TagName  string // By default is `@ctor`
		Template string // By default defined in defaultCtorTemplate
		Target   string // By default is `cfg_annotated.go`
	}
	// CtorTmplData template
	CtorTmplData struct {
		Signature typgen.Signature
		Package   string
		Imports   map[string]string
		Ctors     []*Ctor
	}
	// Ctor is constructor model
	Ctor struct {
		Name string `json:"name"`
		Def  string `json:"-"`
	}
)

const defaultCtorTemplate = `package {{.Package}}

/* {{.Signature}}*/

import ({{range $import, $name := .Imports}}
	{{$name}} "{{$import}}"{{end}}
)

func init() { {{if .Ctors}}{{range $c := .Ctors}}
	typapp.Provide("{{$c.Name}}", {{$c.Def}}){{end}}{{end}}
}`

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
	if a.Template == "" {
		a.Template = defaultCtorTemplate
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
	var ctors []*Ctor
	importAliases := typgen.NewImportAliases()
	for _, directive := range directives {
		alias := importAliases.Append(directive.Package())
		ctors = append(ctors, &Ctor{
			Name: directive.TagParam.Get("name"),
			Def:  fmt.Sprintf("%s.%s", alias, directive.GetName()),
		})

	}
	importAliases.Map["github.com/typical-go/typical-go/pkg/typapp"] = ""

	dest := filepath.Dir(a.Target)
	os.MkdirAll(dest, 0777)
	c.Infof("Generate @ctor to %s\n", a.Target)
	err := tmplkit.WriteFile(a.Target, a.Template, &CtorTmplData{
		Signature: typgen.Signature{TagName: a.TagName},
		Package:   filepath.Base(dest),
		Imports:   importAliases.Map,
		Ctors:     ctors,
	})

	if err != nil {
		return err
	}
	typgo.GoImports(c, a.Target)
	return nil
}

//
// Ctor
//

func (c Ctor) String() string {
	return fmt.Sprintf("{Name=%s Def=%s}", c.Name, c.Def)
}
