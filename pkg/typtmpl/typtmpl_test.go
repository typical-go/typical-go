package typtmpl_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestExecute(t *testing.T) {
	require.EqualError(t,
		typtmpl.Execute("", "bad-template {{{}", nil, nil),
		"template: :1: unexpected \"{\" in command",
	)

	var debugger strings.Builder
	require.NoError(t,
		typtmpl.Execute("", "hello {{.Name}}", &data{Name: "world"}, &debugger),
	)
	require.Equal(t, "hello world", debugger.String())
}

type data struct {
	Name string
}

type dummyTemplate struct {
	text string
}

func (s *dummyTemplate) Execute(w io.Writer) (err error) {
	w.Write([]byte(s.text))
	return
}

type testcase struct {
	testName string
	typtmpl.Template
	expected string
}

func testTemplate(t *testing.T, cases ...testcase) {
	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {
			var debugger strings.Builder
			require.NoError(t, tt.Execute(&debugger))
			require.Equal(t, tt.expected, debugger.String())
		})
	}
}
