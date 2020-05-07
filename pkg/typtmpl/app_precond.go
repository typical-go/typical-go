package typtmpl

import (
	"io"
	"reflect"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

var _ Template = (*AppPrecond)(nil)

const appPrecond = `typapp.Provide({{range $c := .Ctors}}
	&typapp.Constructor{
		Name: "{{$c.Name}}", 
		Fn: {{$c.Def}},
	},{{end}}{{range $c := .CfgCtors}}
	&typapp.Constructor{
		Name: "{{$c.Name}}", 
		Fn: func() (cfg {{$c.SpecType}}, err error) {
			cfg = new({{$c.SpecType2}})
			if err = typcfg.Process("{{$c.Prefix}}", cfg); err != nil {
				return nil, err
			}
			return
		},
	},{{end}}
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

// NewAppPrecond return new instance of ProvideCtor
func NewAppPrecond() *AppPrecond {
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
func (t *AppPrecond) AppendCfgCtor(cfg *typcfg.Configuration) {
	specType := reflect.TypeOf(cfg.Spec).String()
	t.CfgCtors = append(t.CfgCtors, &CfgCtor{
		Name:      cfg.CtorName,
		Prefix:    cfg.Name,
		SpecType:  specType,
		SpecType2: specType[1:],
	})
}

// Execute app precondition template
func (t *AppPrecond) Execute(w io.Writer) (err error) {
	return Execute("appPrecond", appPrecond, t, w)
}
