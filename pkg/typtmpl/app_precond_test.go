package typtmpl_test

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			expected: `typgo.Provide(
	&typgo.Constructor{Name: "", Fn: pkg1.NewFunction1},
	&typgo.Constructor{Name: "", Fn: pkg2.NewFunction2},
)
typgo.Destroy(
)`,
		},
		testcase{
			testName: "constructor for configuration",
			Template: &typtmpl.AppPrecond{
				CfgCtors: []*typtmpl.CfgCtor{
					{Name: "", Prefix: "AAA", SpecType: "*Sample", SpecType2: "Sample"},
				},
			},
			expected: `typgo.Provide(
	&typgo.Constructor{
		Name: "", 
		Fn: func() (cfg *Sample, err error) {
			cfg = new(Sample)
			if err = typgo.Process("AAA", cfg); err != nil {
				return nil, err
			}
			return
		},
	},
)
typgo.Destroy(
)`,
		},
		testcase{
			testName: "constructor for configuration",
			Template: &typtmpl.AppPrecond{
				Dtors: []*typtmpl.Dtor{
					{Def: "pkg1.NewFunction1"},
				},
			},
			expected: `typgo.Provide(
)
typgo.Destroy(
	&typgo.Destructor{Fn: pkg1.NewFunction1},
)`,
		},
	)
}

func TestAppPrecond_NotEmpty(t *testing.T) {
	testcases := []struct {
		testname string
		typtmpl.AppPrecond
		expected bool
	}{
		{
			AppPrecond: typtmpl.AppPrecond{
				Ctors:    []*typtmpl.Ctor{},
				CfgCtors: []*typtmpl.CfgCtor{},
				Dtors:    []*typtmpl.Dtor{},
			},
			expected: false,
		},
		{
			AppPrecond: typtmpl.AppPrecond{
				Ctors: []*typtmpl.Ctor{
					&typtmpl.Ctor{},
				},
				CfgCtors: []*typtmpl.CfgCtor{},
				Dtors:    []*typtmpl.Dtor{},
			},
			expected: true,
		},
		{
			AppPrecond: typtmpl.AppPrecond{
				Ctors: []*typtmpl.Ctor{},
				CfgCtors: []*typtmpl.CfgCtor{
					&typtmpl.CfgCtor{},
				},
				Dtors: []*typtmpl.Dtor{},
			},
			expected: true,
		},
		{
			AppPrecond: typtmpl.AppPrecond{
				Ctors:    []*typtmpl.Ctor{},
				CfgCtors: []*typtmpl.CfgCtor{},
				Dtors: []*typtmpl.Dtor{
					&typtmpl.Dtor{},
				},
			},
			expected: true,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testname, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.NotEmpty())
		})
	}
}

type sample struct{}
