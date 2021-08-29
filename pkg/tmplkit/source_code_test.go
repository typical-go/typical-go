package tmplkit_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/tmplkit"
)

func TestStringer(t *testing.T) {
	testnames := []struct {
		TestName string
		Stringer fmt.Stringer
		Expected string
	}{
		{
			TestName: "empty signature",
			Stringer: tmplkit.Signature{},
			Expected: "/* DO NOT EDIT. This is code generated file. */",
		},
		{
			TestName: "signature with tag",
			Stringer: tmplkit.Signature{TagName: "@some-tag"},
			Expected: "/* DO NOT EDIT. This is code generated file from '@some-tag' annotation. */",
		},
		{
			TestName: "line code",
			Stringer: tmplkit.LineCode(`fmt.Println("Hello World")`),
			Expected: `fmt.Println("Hello World")`,
		},
		{
			TestName: "formatted line code",
			Stringer: tmplkit.LineCodef(`fmt.Println("%s")`, "abcdefgh"),
			Expected: `fmt.Println("abcdefgh")`,
		},
		{
			TestName: "comment",
			Stringer: tmplkit.Comment(`some comment`),
			Expected: `// some comment`,
		},
		{
			TestName: "init function",
			Stringer: tmplkit.InitFunction([]fmt.Stringer{
				tmplkit.LineCode(`fmt.Println("Hello World")`),
			}),
			Expected: "func init(){\n\tfmt.Println(\"Hello World\")\n}\n",
		},
		{
			TestName: "function",
			Stringer: &tmplkit.Function{
				Name: "SomeFunction",
				Params: []tmplkit.Variable{
					{Name: "a", Type: "string"},
					{Name: "b", Type: "int64"},
				},
				Returns: []tmplkit.Variable{
					{Name: "s", Type: "string"},
					{Name: "err", Type: "error"},
				},
				LineCodes: []fmt.Stringer{
					tmplkit.LineCode(`fmt.Println("Hello World")`),
				},
			},
			Expected: "func SomeFunction(a string,b int64)(s string,err error){\n\tfmt.Println(\"Hello World\")\n}\n",
		},
		{
			TestName: "source code",
			Stringer: &tmplkit.SourceCode{
				Package: "utils",
				Imports: tmplkit.NewImports(map[string]string{
					"fmt":               "",
					"github.com/lib/pq": "_",
				}),
				LineCodes: []fmt.Stringer{
					tmplkit.Comment("some comment"),
				},
			},
			Expected: "package utils\n\n/* DO NOT EDIT. This is code generated file. */\nimport (\n\t\"fmt\"\n\t_ \"github.com/lib/pq\"\n)\n\n// some comment\n",
		},
		{
			TestName: "imports",
			Stringer: tmplkit.NewImports(map[string]string{
				"fmt":               "",
				"github.com/lib/pq": "_",
			}),
			Expected: "import (\n\t\"fmt\"\n\t_ \"github.com/lib/pq\"\n)\n\n",
		},
		{
			TestName: "empty imports",
			Stringer: tmplkit.NewImports(nil),
			Expected: "",
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Stringer.String())
		})
	}
}
