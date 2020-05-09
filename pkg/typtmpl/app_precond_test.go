package typtmpl_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestProvideCtor(t *testing.T) {
	testTemplate(t,
		testcase{
			testName: "common constructor",
			Template: &typtmpl.AppPrecond{
				Ctors: []*typtmpl.Ctor{
					{Name: "", Def: "pkg1.NewFunction1"},
					{Name: "", Def: "pkg2.NewFunction2"},
				},
			},
			expected: `typapp.Provide(
	&typapp.Constructor{Name: "", Fn: pkg1.NewFunction1},
	&typapp.Constructor{Name: "", Fn: pkg2.NewFunction2},
)
typapp.Destory(
)`,
		},
		testcase{
			testName: "constructor for configuration",
			Template: &typtmpl.AppPrecond{
				CfgCtors: []*typtmpl.CfgCtor{
					{Name: "", Prefix: "AAA", SpecType: "*Sample", SpecType2: "Sample"},
				},
			},
			expected: `typapp.Provide(
	&typapp.Constructor{
		Name: "", 
		Fn: func() (cfg *Sample, err error) {
			cfg = new(Sample)
			if err = typcfg.Process("AAA", cfg); err != nil {
				return nil, err
			}
			return
		},
	},
)
typapp.Destory(
)`,
		},
		testcase{
			testName: "constructor for configuration",
			Template: &typtmpl.AppPrecond{
				Dtors: []*typtmpl.Dtor{
					{Def: "pkg1.NewFunction1"},
				},
			},
			expected: `typapp.Provide(
)
typapp.Destory(
	&typapp.Destructor{Fn: pkg1.NewFunction1},
)`,
		},
	)
}

type sample struct{}
