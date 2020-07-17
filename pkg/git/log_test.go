package git_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/git"
)

func TestLog(t *testing.T) {
	testcases := []struct {
		TestName string
		Raw      string
		Expected *git.Log
	}{
		{
			Raw: "123456",
		},
		{
			Raw:      "5378feb",
			Expected: &git.Log{ShortCode: "5378feb"},
		},
		{
			Raw:      "5378feb something",
			Expected: &git.Log{ShortCode: "5378feb", Message: "something"},
		},
		{
			Raw:      "5378feb one two three four",
			Expected: &git.Log{ShortCode: "5378feb", Message: "one two three four"},
		},
		{
			Raw:      "5378feb something \n\nCo-Authored-By: xx <xx@users.noreply.github.com>",
			Expected: &git.Log{ShortCode: "5378feb", Message: "something", CoAuthoredBy: "xx <xx@users.noreply.github.com>"},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.EqualValues(t, tt.Expected, git.CreateLog(tt.Raw))
		})

	}
}
