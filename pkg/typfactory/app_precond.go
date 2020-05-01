package typfactory

import (
	"io"
	"reflect"
	"text/template"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

const appPrecond = `typapp.Provide({{range $c := .Ctors}}
	typapp.NewConstructor("{{$c.Name}}", {{$c.Def}}),{{end}}{{range $c := .CfgCtors}}
	typapp.NewConstructor("{{$c.Name}}", func() (cfg {{$c.SpecType}}, err error) {
		cfg = new({{$c.SpecType2}})
		if err = typcfg.Process("{{$c.Prefix}}", cfg); err != nil {
			return nil, err
		}
		return
	}),{{end}}
)`

// AppPrecond to generate provide constructor
type AppPrecond struct {
	Ctors    []*Ctor
	CfgCtors []*CfgCtor
}

// Ctor is constructor model
type Ctor struct {
	Name string
	Def  string
}

// CfgCtor is config constructor model
type CfgCtor struct {
	Name      string
	Prefix    string
	SpecType  string
	SpecType2 string
}

// NewProvideCtor return new instance of ProvideCtor
func NewProvideCtor() *AppPrecond {
	return &AppPrecond{}
}

// AppendCtor to append constructor
func (t *AppPrecond) AppendCtor(name, def string) {
	t.Ctors = append(t.Ctors, &Ctor{
		Name: name,
		Def:  def,
	})
}

// AppendCfgCtor to append config constructor
func (t *AppPrecond) AppendCfgCtor(name string, cfg *typcfg.Configuration) {
	specType := reflect.TypeOf(cfg.Spec).String()
	t.CfgCtors = append(t.CfgCtors, &CfgCtor{
		Name:      name,
		Prefix:    cfg.Name,
		SpecType:  specType,
		SpecType2: specType[1:],
	})
}

// Write the tyicalw
func (t *AppPrecond) Write(w io.Writer) (err error) {
	var tmpl *template.Template
	if tmpl, err = template.New("AppPrecond").Parse(appPrecond); err != nil {
		return
	}
	return tmpl.Execute(w, t)
}
