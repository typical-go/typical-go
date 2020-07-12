package typtmpl_test

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestParse(t *testing.T) {
	t.Run("WHEN success", func(t *testing.T) {
		var builder strings.Builder
		require.NoError(t, typtmpl.Parse("", "hello {{.Name}}", &data{Name: "world"}, &builder))
		require.Equal(t, "hello world", builder.String())
	})
	t.Run("WHEN error", func(t *testing.T) {
		require.EqualError(t,
			typtmpl.Parse("", "bad-template {{{}", nil, nil),
			"template: :1: unexpected \"{\" in command",
		)
	})
}

func TestExecuteToFile(t *testing.T) {
	target := "dummy"
	typtmpl.ExecuteToFile(target, &dummyTemplate{"some-parsed-text"})
	defer os.Remove(target)
	b, _ := ioutil.ReadFile(target)
	require.Equal(t, []byte("some-parsed-text"), b)
}

type (
	data struct {
		Name string
	}
	dummyTemplate struct {
		text string
	}
)

func (s *dummyTemplate) Execute(w io.Writer) (err error) {
	w.Write([]byte(s.text))
	return
}
