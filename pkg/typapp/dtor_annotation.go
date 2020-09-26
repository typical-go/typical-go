package typapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/tmplkit"
	"github.com/typical-go/typical-go/pkg/typast"
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
		Signature typast.Signature
		Package   string
		Imports   map[string]string
		Dtors     []*Dtor
	}
	// Dtor is destructor model
	Dtor struct {
		Def string
	}
)

const defaultDtorTemplate = `package {{.Package}}

/* {{.Signature}}*/

import ({{range $import, $name := .Imports}}
	{{$name}} "{{$import}}"{{end}}
)

func init() { {{if .Dtors}}
	typapp.AppendDtor({{range $d := .Dtors}}
		&typapp.Destructor{Fn: {{$d.Def}}},{{end}}
	){{end}}
}`

const (
	dtorHelp = "https://pkg.go.dev/github.com/typical-go/typical-go/pkg/typapp?tab=doc#DtorAnnotation"
)

var _ typast.Annotator = (*DtorAnnotation)(nil)

// Annotate @dtor
func (a *DtorAnnotation) Annotate(c *typast.Context) error {
	annots, imports := FindAnnotFunc(c, a.getTagName())
	imports["github.com/typical-go/typical-go/pkg/typapp"] = ""

	var dtors []*Dtor
	for _, annot := range annots {
		dtors = append(dtors, &Dtor{
			Def: fmt.Sprintf("%s.%s", annot.ImportAlias, annot.GetName()),
		})
	}

	target := fmt.Sprintf("%s/%s", c.Destination, a.getTarget(c))
	if len(dtors) < 1 {
		os.Remove(target)
		return nil
	}

	data := &DtorTmplData{
		Signature: typast.Signature{
			TagName: a.getTagName(),
			Help:    dtorHelp,
		},
		Package: filepath.Base(c.Destination),
		Imports: imports,
		Dtors:   dtors,
	}
	fmt.Fprintf(Stdout, "Generate @dtor to %s\n", target)
	if err := tmplkit.WriteFile(target, a.getTemplate(), data); err != nil {
		return err
	}
	typgo.GoImports(target)
	return nil
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

func (a *DtorAnnotation) getTarget(c *typast.Context) string {
	if a.Target == "" {
		a.Target = "dtor_annotated.go"
	}
	return a.Target
}
