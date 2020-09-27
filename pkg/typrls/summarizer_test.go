package typrls_test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestSummarizer(t *testing.T) {
	testCases := []struct {
		TestName string
		typrls.Summarizer
		Context         *typgo.Context
		RunExpectations []*execkit.RunExpectation
		Expected        string
		ExpectedOut     string
	}{
		{
			Summarizer: typrls.NewSummarizer(func(*typgo.Context) string {
				return "some-text"
			}),
			Expected: "some-text",
		},
		{
			TestName: "change summary",
			Summarizer: &typrls.GitSummarizer{
				ExcludePrefix: []string{"merge", "revision"},
			},
			Context: &typgo.Context{
				Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
			},
			RunExpectations: []*execkit.RunExpectation{
				{
					CommandLine: "git describe --tags --abbrev=0",
					OutputBytes: []byte("v0.0.1"),
				},
				{
					CommandLine: "git --no-pager log v0.0.1..HEAD --oneline",
					OutputBytes: []byte("1234567 some-message-1\n1234568 some-message-3"),
				},
			},
			Expected:    "1234567 some-message-1\n1234568 some-message-3",
			ExpectedOut: "\n$ git describe --tags --abbrev=0\n\n$ git --no-pager log v0.0.1..HEAD --oneline\n",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			var out strings.Builder
			typgo.Stdout = &out
			defer func() { typgo.Stdout = os.Stdout }()

			unpatch := execkit.Patch(tt.RunExpectations)
			defer unpatch(t)
			require.Equal(t, tt.Expected, tt.Summarize(tt.Context))
			require.Equal(t, tt.ExpectedOut, out.String())
		})
	}
}

func TestChangeSummarize_HasPrefix(t *testing.T) {
	summarizer := &typrls.GitSummarizer{
		ExcludePrefix: []string{"merge", "revision"},
	}

	testcases := []struct {
		TestName string
		Msg      string
		Expected bool
	}{
		{Msg: "Merge something", Expected: true},
		{Msg: "merge something", Expected: true},
		{Msg: "MERGE something", Expected: true},
		{Msg: "revision: something", Expected: true},
		{Msg: "asdf", Expected: false},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, summarizer.HasPrefix(tt.Msg))
		})
	}
}
