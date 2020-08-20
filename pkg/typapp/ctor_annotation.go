package typapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// CtorAnnotation handle @ctor annotation
	// e.g. `@ctor (name:"NAME")`
	CtorAnnotation struct {
		TagName  string // By default is `@ctor`
		Template string // By default defined in defaultCtorTemplate
		Target   string // By default is `cfg_annotated.go`
	}
	// CtorTmplData template
	CtorTmplData struct {
		Signature typannot.Signature
		Package   string
		Imports   []string
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

import ({{range $import := .Imports}}
	"{{$import}}"{{end}}
)

func init() { {{if .Ctors}}
	typapp.AppendCtor({{range $c := .Ctors}}
		&typapp.Constructor{Name: "{{$c.Name}}", Fn: {{$c.Def}}},{{end}}
	){{end}}
}`

const ctorHelp = "https://pkg.go.dev/github.com/typical-go/typical-go/pkg/typapp?tab=doc#CtorAnnotation"

var _ typannot.Annotator = (*CtorAnnotation)(nil)

// Annotate ctor
func (a *CtorAnnotation) Annotate(c *typannot.Context) error {
	ctors := a.CreateCtors(c)
	target := fmt.Sprintf("%s/%s", c.Destination, a.getTarget(c))
	pkg := filepath.Base(c.Destination)

	if len(ctors) < 1 {
		os.Remove(target)
		return nil
	}

	data := &CtorTmplData{
		Signature: typannot.Signature{
			TagName: a.getTagName(),
			Help:    ctorHelp,
		},
		Package: pkg,
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typapp",
		),
		Ctors: ctors,
	}
	fmt.Fprintf(Stdout, "Generate @ctor to %s\n", target)
	if err := common.ExecuteTmplToFile(target, a.getTemplate(), data); err != nil {
		return err
	}
	typgo.GoImports(target)
	return nil
}

// CreateCtors get ctors
func (a *CtorAnnotation) CreateCtors(c *typannot.Context) []*Ctor {
	var ctors []*Ctor
	for _, annot := range c.FindAnnotByFunc(a.getTagName()) {
		ctors = append(ctors, &Ctor{
			Name: annot.TagParam.Get("name"),
			Def:  fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name),
		})
	}
	return ctors
}

func (a *CtorAnnotation) getTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = "ctor_annotated.go"
	}
	return a.Target
}

func (a *CtorAnnotation) getTagName() string {
	if a.TagName == "" {
		a.TagName = "@ctor"
	}
	return a.TagName
}

func (a *CtorAnnotation) getTemplate() string {
	if a.Template == "" {
		a.Template = defaultCtorTemplate
	}
	return a.Template
}
