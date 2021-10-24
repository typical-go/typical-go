package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestFunction(t *testing.T) {
	func1 := &typgen.Function{
		Name: "some-name",
		Docs: []string{"doc1", "doc2"},
	}
	require.Equal(t, "some-name", func1.GetName())
	require.Equal(t, []string{"doc1", "doc2"}, func1.GetDocs())
}

func TestFunction_IsMethod(t *testing.T) {
	testnames := []struct {
		TestName string
		*typgen.Function
		Expected bool
	}{
		{
			Function: &typgen.Function{Recv: []*typgen.Field{{}}},
			Expected: true,
		},
		{
			Function: &typgen.Function{},
			Expected: false,
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.IsMethod())
		})
	}
}

func TestFunction_SourceCode(t *testing.T) {
	testCases := []struct {
		TestName string
		Function *typgen.Function
		Expected string
	}{
		{
			Function: &typgen.Function{
				Name: "someFunc",
			},
			Expected: `func someFunc(){
}`,
		},
		{
			Function: &typgen.Function{
				Name: "someFunc",
				Params: []*typgen.Field{
					{Names: []string{"arg1", "arg2"}, Type: "string"},
					{Names: []string{"arg3"}, Type: "int64"},
				},
				Body: typgen.CodeLine(`fmt.Println("hello world")`),
				Returns: []*typgen.Field{
					{Type: "string"},
					{Type: "error"},
				},
			},
			Expected: `func someFunc(arg1,arg2 string,arg3 int64)( string, error){
fmt.Println("hello world")
}`,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Function.Code())
		})
	}
}
