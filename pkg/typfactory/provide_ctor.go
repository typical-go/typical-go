package typfactory

import (
	"io"
	"reflect"
	"text/template"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

const provideCtor = `typapp.AppendConstructor({{range $def := .FnDefs}}
	typapp.NewConstructor({{$def}}),{{end}}{{range $c := .Cfgs}}
	typapp.NewConstructor(func() (cfg {{$.SpecType $c}}, err error) {
		cfg = new({{$.SpecType2 $c}})
		if err = typcfg.Process("{{$c.Name}}", cfg); err != nil {
			return nil, err
		}
		return
	}),{{end}}
)`

// ProvideCtor to generate provide constructor
type ProvideCtor struct {
	FnDefs []string
	Cfgs   []*typcfg.Configuration
}

// Write the tyicalw
func (t *ProvideCtor) Write(w io.Writer) (err error) {
	var tmpl *template.Template
	if tmpl, err = template.New("ProvideCtor").Parse(provideCtor); err != nil {
		return
	}
	return tmpl.Execute(w, t)
}

func (t *ProvideCtor) SpecType(c *typcfg.Configuration) string {
	return reflect.TypeOf(c.Spec).String()
}

func (t *ProvideCtor) SpecType2(c *typcfg.Configuration) string {
	return t.SpecType(c)[1:]
}
