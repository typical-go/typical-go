package typtmpl

import (
	"io"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

type (
	// Template responsible to write
	Template interface {
		Execute(io.Writer) error
	}
	// TestCase for template
	TestCase struct {
		TestName string
		Template
		Expected string
	}
)

// Execute template
func Execute(name, text string, data interface{}, w io.Writer) (err error) {
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return
	}
	return tmpl.Execute(w, data)
}

// TestTemplate to test template
func TestTemplate(t *testing.T, cases []TestCase) {
	for _, tt := range cases {
		t.Run(tt.TestName, func(t *testing.T) {
			var debugger strings.Builder
			require.NoError(t, tt.Execute(&debugger))
			require.Equal(t, tt.Expected, debugger.String())
		})
	}
}
