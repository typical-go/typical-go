package common_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/common"
)

func TestExecuteTmpl(t *testing.T) {
	var out strings.Builder
	require.NoError(t, common.ExecuteTmpl(&out, "hello {{.Name}}", &data{Name: "world"}))
	require.Equal(t, "hello world", out.String())
}

func TestExecuteTmpl_Error(t *testing.T) {
	require.EqualError(t,
		common.ExecuteTmpl(nil, "bad-template {{{}", &struct{}{}),
		"template: *struct {}:1: unexpected \"{\" in command",
	)
}

func TestExecuteTmplToFile(t *testing.T) {
	target := "sample-target"
	defer os.Remove(target)
	require.NoError(t, common.ExecuteTmplToFile(target, "hello {{.Name}}", &data{Name: "world"}))

	b, _ := ioutil.ReadFile(target)
	require.Equal(t, "hello world", string(b))
}

type (
	data struct {
		Name string
	}
)
