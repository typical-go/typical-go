package stdrelease_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/stdrelease"
)

func TestStandardFilter(t *testing.T) {
	testcases := []struct {
		ignorings []string
		messages  []string
		filtereds []string
	}{
		{
			[]string{"revision"},
			[]string{"5378feb revision: something"},
			[]string{},
		},
		{
			[]string{"revision"},
			[]string{"5378feb REVISION: something"},
			[]string{},
		},
		{
			[]string{},
			[]string{"5378feb something \n\nCo-Authored-By: xx <xx@users.noreply.github.com>"},
			[]string{"5378feb something"},
		},
	}
	for _, tt := range testcases {
		filter := stdrelease.StandardFilter{tt.ignorings}
		require.Equal(t, tt.filtereds, filter.Filter(tt.messages))
	}
}

func TestCleanMessage(t *testing.T) {
	testcases := []struct {
		message string
		cleaned string
	}{
		{message: "    abcde    \n", cleaned: "abcde"},
		{message: "some message\n\nCo-Authored-By: xx <xx@users.noreply.github.com>", cleaned: "some message"},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.cleaned, stdrelease.CleanMessage(tt.message))
	}
}

func TestMessage(t *testing.T) {
	testcases := []struct {
		changelog string
		message   string
	}{
		{
			"5378feb rename versioning to tagging and combine goos and goarch as target",
			"rename versioning to tagging and combine goos and goarch as target",
		},
		{"5378feb ", ""},
		{"5378feb", ""},
		{"", ""},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.message, stdrelease.MessageText(tt.changelog))
	}
}
