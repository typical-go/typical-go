package typrls_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typrls"
)

func TestSummarizer(t *testing.T) {
	testCases := []struct {
		TestName string
		typrls.Summarizer
		*typrls.Context
		Expected string
	}{
		{
			Summarizer: typrls.NewSummarizer(func(*typrls.Context) string {
				return "some-text"
			}),
			Expected: "some-text",
		},
		{
			TestName: "change summary",
			Summarizer: &typrls.GitSummarizer{
				ExcludePrefix: []string{"merge", "revision"},
			},
			Context: &typrls.Context{
				Git: &typrls.Git{
					Logs: []*typrls.Log{
						{ShortCode: "1111", Message: "some-message-1"},
						{ShortCode: "2222", Message: "revision: some-message-2"},
						{ShortCode: "3333", Message: "some-message-3"},
					},
				},
			},
			Expected: "1111 some-message-1\n3333 some-message-3",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Summarize(tt.Context))
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
