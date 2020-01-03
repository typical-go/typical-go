package golang_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/golang"
)

func TestFunc(t *testing.T) {
	t.Run("with param & return", func(t *testing.T) {
		fn := golang.Function{
			Name: "myFunction",
			Params: []*common.StringKV{
				{Key: "num", Value: "int"},
				{Key: "text", Value: "string"},
			},
			Returns: []*common.StringKV{
				{Key: "altered", Value: "string"},
				{Key: "err", Value: "error"},
			},
		}
		fn.Append(`fmt.Println("Hello World")`)
		fn.Return("err")
		var b strings.Builder
		fn.Write(&b)
		require.Equal(t, "func myFunction(num int, text string) (altered string, err error) {\nfmt.Println(\"Hello World\")\nreturn err\n}\n", b.String())
	})
	t.Run("with param & no return", func(t *testing.T) {
		fn := golang.Function{
			Name: "myFunction",
			Params: []*common.StringKV{
				{Key: "num", Value: "int"},
			},
		}
		fn.Append(`fmt.Println("Hello World")`)
		var b strings.Builder
		fn.Write(&b)
		require.Equal(t, "func myFunction(num int) {\nfmt.Println(\"Hello World\")\n}\n", b.String())
	})
}
