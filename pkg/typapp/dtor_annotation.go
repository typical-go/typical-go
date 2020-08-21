package typapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/tmplkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// DtorAnnotation handle @dtor annotation. No Attributes required.
	DtorAnnotation struct {
		TagName  string // By default is @dtor
		Template string // By default defined in defaultDtorTemplate
		Target   string // By default is `cmd/PROJECT_NAME/dtor_annotated.go`
	}
	// DtorTmplData template
	DtorTmplData struct {
		Signature typannot.Signature
		Package   string
		Imports   []string
		Dtors     []*Dtor
	}
	// Dtor is destructor model
	Dtor struct {
		Def string
	}
)

const defaultDtorTemplate = `package {{.Package}}

/* {{.Signature}}*/

import ({{range $import := .Imports}}
	"{{$import}}"{{end}}
)

func init() { {{if .Dtors}}
	typapp.AppendDtor({{range $d := .Dtors}}
		&typapp.Destructor{Fn: {{$d.Def}}},{{end}}
	){{end}}
}`

const dtorHelp = "https://pkg.go.dev/github.com/typical-go/typical-go/pkg/typapp?tab=doc#DtorAnnotation"

var _ typannot.Annotator = (*DtorAnnotation)(nil)

// Annotate @dtor
func (a *DtorAnnotation) Annotate(c *typannot.Context) error {
	dtors := a.CreateDtors(c)
	target := fmt.Sprintf("%s/%s", c.Destination, a.getTarget(c))
	pkg := filepath.Base(c.Destination)
	if len(dtors) < 1 {
		os.Remove(target)
		return nil
	}
	data := &DtorTmplData{
		Signature: typannot.Signature{
			TagName: a.getTagName(),
			Help:    dtorHelp,
		},
		Package: pkg,
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typapp",
		),
		Dtors: dtors,
	}
	fmt.Fprintf(Stdout, "Generate @dtor to %s\n", target)
	if err := tmplkit.WriteFile(target, a.getTemplate(), data); err != nil {
		return err
	}
	typgo.GoImports(target)
	return nil
}

// CreateDtors get dtors
func (a *DtorAnnotation) CreateDtors(c *typannot.Context) []*Dtor {
	var dtors []*Dtor
	for _, annot := range c.FindAnnotByFunc(a.getTagName()) {
		dtors = append(dtors, &Dtor{
			Def: fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name),
		})
	}
	return dtors
}

func (a *DtorAnnotation) getTagName() string {
	if a.TagName == "" {
		a.TagName = "@dtor"
	}
	return a.TagName
}

func (a *DtorAnnotation) getTemplate() string {
	if a.Template == "" {
		a.Template = defaultDtorTemplate
	}
	return a.Template
}

func (a *DtorAnnotation) getTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = "dtor_annotated.go"
	}
	return a.Target
}
