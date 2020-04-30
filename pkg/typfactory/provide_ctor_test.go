package typfactory_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

func TestProvideCtor(t *testing.T) {
	testcases := []testcase{
		{
			testName: "common constructor",
			Writer: &typfactory.ProvideCtor{
				Ctors: []*typfactory.Ctor{
					{Name: "", Def: "pkg1.NewFunction1"},
					{Name: "", Def: "pkg2.NewFunction2"},
				},
			},
			expected: `typapp.AppendConstructor(
	typapp.NewConstructor("", pkg1.NewFunction1),
	typapp.NewConstructor("", pkg2.NewFunction2),
)`,
		},
		{
			testName: "constructor for configuration",
			Writer: &typfactory.ProvideCtor{
				CfgCtors: []*typfactory.CfgCtor{
					{Name: "", Prefix: "AAA", SpecType: "*Sample", SpecType2: "Sample"},
				},
			},
			expected: `typapp.AppendConstructor(
	typapp.NewConstructor("", func() (cfg *Sample, err error) {
		cfg = new(Sample)
		if err = typcfg.Process("AAA", cfg); err != nil {
			return nil, err
		}
		return
	}),
)`,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			var debugger strings.Builder
			require.NoError(t, tt.Write(&debugger))
			require.Equal(t, tt.expected, debugger.String())
		})

	}
}

type sample struct{}
