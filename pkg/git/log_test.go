package git_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/git"
)

func TestLog(t *testing.T) {
	testcases := []struct {
		raw      string
		expected *git.Log
	}{
		{raw: "123456"},
		{raw: "5378feb", expected: &git.Log{Short: "5378feb"}},
		{raw: "5378feb something", expected: &git.Log{Short: "5378feb", Message: "something"}},
		{raw: "5378feb one two three four", expected: &git.Log{Short: "5378feb", Message: "one two three four"}},
		{raw: "5378feb something \n\nCo-Authored-By: xx <xx@users.noreply.github.com>", expected: &git.Log{Short: "5378feb", Message: "something", CoAuthoredBy: "xx <xx@users.noreply.github.com>"}},
	}
	for _, tt := range testcases {
		require.EqualValues(t, tt.expected, git.CreateLog(tt.raw))
	}
}
