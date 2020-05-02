package typtmpl_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestGitIgnore(t *testing.T) {
	testTemplate(t,
		testcase{
			Template: &typtmpl.GitIgnore{},
			expected: `/bin
/release
/.typical-tmp
/vendor 
*.envrc
*.env
*.test
*.out
`,
		},
	)

}
