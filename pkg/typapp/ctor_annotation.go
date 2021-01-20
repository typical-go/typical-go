package typapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/oskit"
	"github.com/typical-go/typical-go/pkg/tmplkit"
	"github.com/typical-go/typical-go/pkg/typast"
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
		Signature typast.Signature
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
// CtorAnnotation
//

var _ typast.Annotator = (*CtorAnnotation)(nil)

// Annotate ctor
func (a *CtorAnnotation) Annotate(c *typast.Context) error {
	annots, imports := typast.FindAnnot(c, a.getTagName(), typast.EqualFunc)
	imports["github.com/typical-go/typical-go/pkg/typapp"] = ""

	var ctors []*Ctor
	for _, annot := range annots {
		ctors = append(ctors, &Ctor{
			Name: annot.TagParam.Get("name"),
			Def:  fmt.Sprintf("%s.%s", annot.ImportAlias, annot.GetName()),
		})
	}

	target := a.getTarget(c)
	dest := filepath.Dir(target)
	os.MkdirAll(dest, 0777)

	data := &CtorTmplData{
		Signature: typast.Signature{TagName: a.getTagName()},
		Package:   filepath.Base(dest),
		Imports:   imports,
		Ctors:     ctors,
	}

	fmt.Fprintf(oskit.Stdout, "Generate @ctor to %s\n", target)
	if err := tmplkit.WriteFile(target, a.getTemplate(), data); err != nil {
		return err
	}
	typgo.GoImports(target)
	return nil
}

func (a *CtorAnnotation) getTarget(c *typast.Context) string {
	if a.Target == "" {
		a.Target = "internal/generated/ctor/ctor.go"
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

//
// Ctor
//

func (c Ctor) String() string {
	return fmt.Sprintf("{Name=%s Def=%s}", c.Name, c.Def)
}
