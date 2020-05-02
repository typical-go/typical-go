package typtmpl_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
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
	typapp.NewConstructor("", pkg1.NewFunction1),
	typapp.NewConstructor("", pkg2.NewFunction2),
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
	typapp.NewConstructor("", func() (cfg *Sample, err error) {
		cfg = new(Sample)
		if err = typcfg.Process("AAA", cfg); err != nil {
			return nil, err
		}
		return
	}),
)`,
		},
	)
}

func TestAppend(t *testing.T) {
	appPrecond := typtmpl.NewAppPrecond()
	appPrecond.AppendCtor("some-name", "some-def")
	appPrecond.AppendCtor("some-name2", "some-def2")
	appPrecond.AppendCfgCtor(&typcfg.Configuration{
		Name: "prefix",
		Spec: &sample{},
	})
	appPrecond.AppendCfgCtor(&typcfg.Configuration{
		CtorName: "prefix2",
		Name:     "prefix2",
		Spec:     &sample{},
	})

	require.Equal(t, &typtmpl.AppPrecond{
		Ctors: []*typtmpl.Ctor{
			{"some-name", "some-def"},
			{"some-name2", "some-def2"},
		},
		CfgCtors: []*typtmpl.CfgCtor{
			{
				Name:      "",
				Prefix:    "prefix",
				SpecType:  "*typtmpl_test.sample",
				SpecType2: "typtmpl_test.sample",
			},
			{
				Name:      "prefix2",
				Prefix:    "prefix2",
				SpecType:  "*typtmpl_test.sample",
				SpecType2: "typtmpl_test.sample",
			},
		},
	}, appPrecond)
}

type sample struct{}
