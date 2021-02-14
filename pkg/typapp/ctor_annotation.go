package typapp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	if a.TagName == "" {
		a.TagName = "@ctor"
	}
	if a.Target == "" {
		a.Target = "internal/generated/ctor/ctor.go"
	}
	if a.Template == "" {
		a.Template = defaultCtorTemplate
	}

	var ctors []*Ctor
	importAliases := typast.NewImportAliases()
	for _, annot := range c.Annots {
		if strings.EqualFold(annot.TagName, a.TagName) && typast.IsFunc(annot) && typast.IsPublic(annot) {
			alias := importAliases.Append(typast.Package(annot))
			ctors = append(ctors, &Ctor{
				Name: annot.TagParam.Get("name"),
				Def:  fmt.Sprintf("%s.%s", alias, annot.GetName()),
			})
		}
	}
	importAliases.Map["github.com/typical-go/typical-go/pkg/typapp"] = ""

	dest := filepath.Dir(a.Target)
	os.MkdirAll(dest, 0777)
	c.Infof("Generate @ctor to %s\n", a.Target)
	err := tmplkit.WriteFile(a.Target, a.Template, &CtorTmplData{
		Signature: typast.Signature{TagName: a.TagName},
		Package:   filepath.Base(dest),
		Imports:   importAliases.Map,
		Ctors:     ctors,
	})

	if err != nil {
		return err
	}
	typgo.GoImports(c.Context, a.Target)

	return nil
}

//
// Ctor
//

func (c Ctor) String() string {
	return fmt.Sprintf("{Name=%s Def=%s}", c.Name, c.Def)
}
