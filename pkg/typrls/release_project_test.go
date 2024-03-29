package typrls_test

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/urfave/cli/v2"
)

func TestReleaseProject(t *testing.T) {
	defer os.Remove("release")
	var debug strings.Builder
	testcases := []struct {
		TestName string
		typrls.ReleaseProject
		Context     *typgo.Context
		Debug       string
		ExpectedErr string
	}{
		{
			ReleaseProject: typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: typrls.DefaultGenerateSummary,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					fmt.Fprintln(&debug, "1")
					return nil
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					fmt.Fprintln(&debug, "2")
					return nil
				}),
			},
			Context: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			Debug: "1\n2\n",
		},
		{
			TestName: "release error",
			ReleaseProject: typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: typrls.DefaultGenerateSummary,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return errors.New("release-error")
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			ExpectedErr: "release-error",
		},
		{
			TestName: "publish error",
			ReleaseProject: typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: typrls.DefaultGenerateSummary,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return nil
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			ExpectedErr: "publish-error",
		},
		{
			TestName: "empty publisher",
			ReleaseProject: typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: typrls.DefaultGenerateSummary,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return nil
				}),
			},
			Context: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
		},
		{
			TestName: "skip publish",
			ReleaseProject: typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: typrls.DefaultGenerateSummary,
				Releaser: typrls.NewReleaser(func(c *typrls.Context) error {
					return nil
				}),
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:    createContext("-skip-publish"),
				Descriptor: &typgo.Descriptor{},
			},
		},
		{
			TestName: "empty releaser",
			ReleaseProject: typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: typrls.DefaultGenerateSummary,
				Publisher: typrls.NewPublisher(func(c *typrls.Context) error {
					return errors.New("publish-error")
				}),
			},
			Context: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			ExpectedErr: "publish-error",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			defer debug.Reset()
			defer tt.Context.PatchBash([]*typgo.MockCommand{})(t)

			err := tt.Execute(tt.Context)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.Debug, debug.String())
			}
		})
	}

}

func TestReleaseProject_CustomReleaseFolder(t *testing.T) {
	var rlsCtx *typrls.Context

	rel := &typrls.ReleaseProject{
		GenerateTagFn:     typrls.DefaultGenerateTag,
		GenerateSummaryFn: func(*typgo.Context) string { return "some-summary" },
		Releaser: typrls.NewReleaser(func(r *typrls.Context) error {
			rlsCtx = r
			return nil
		}),
	}

	c := &typgo.Context{
		Context:    createContext("-release-folder=some-release"),
		Descriptor: &typgo.Descriptor{ProjectVersion: "9.9.9"},
	}
	defer c.PatchBash(nil)(t)
	rel.Execute(c)
	defer os.RemoveAll("some-release")
	require.Equal(t, "some-release", rlsCtx.ReleaseFolder)
}

func TestReleaseProject_Execute_Context(t *testing.T) {
	defer os.Remove("release")
	testcases := []struct {
		TestName string
		*typrls.ReleaseProject
		Ctx             *typgo.Context
		RunExpectations []*typgo.MockCommand
		Expected        *typrls.Context
		ExpectedErr     string
	}{
		{
			ReleaseProject: &typrls.ReleaseProject{},
			Ctx: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{},
			},
			RunExpectations: []*typgo.MockCommand{
				{CommandLine: "git fetch"},
				{
					CommandLine: "git describe --tags --abbrev=0",
					OutputBytes: []byte("v0.0.1"),
				},
				{
					CommandLine: "git --no-pager log v0.0.1..HEAD --oneline",
					OutputBytes: []byte("1234567 some-message-1\n1234568 some-message-3"),
				},
			},
			Expected: &typrls.Context{
				TagName:       "v0.0.1",
				Alpha:         false,
				Summary:       "1234567 some-message-1\n1234568 some-message-3",
				ReleaseFolder: "release",
			},
		},
		{
			TestName: "with alpha tag",
			ReleaseProject: &typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: func(*typgo.Context) string { return "some-summary" },
			},
			Ctx: &typgo.Context{
				Context:    createContext("-alpha"),
				Descriptor: &typgo.Descriptor{},
			},
			RunExpectations: []*typgo.MockCommand{
				{CommandLine: "git fetch"},
			},
			Expected: &typrls.Context{
				TagName:       "v0.0.1_alpha",
				Alpha:         true,
				Summary:       "some-summary",
				ReleaseFolder: "release",
			},
		},
		{
			TestName: "success",
			ReleaseProject: &typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: func(*typgo.Context) string { return "some-summary" },
			},
			Ctx: &typgo.Context{
				Context:    createContext(),
				Descriptor: &typgo.Descriptor{ProjectVersion: "9.9.9"},
			},
			RunExpectations: []*typgo.MockCommand{
				{CommandLine: "git fetch"},
			},
			Expected: &typrls.Context{
				TagName:       "v9.9.9",
				Alpha:         false,
				Summary:       "some-summary",
				ReleaseFolder: "release",
			},
		},
		{
			TestName: "override release-tag",
			ReleaseProject: &typrls.ReleaseProject{
				GenerateTagFn:     typrls.DefaultGenerateTag,
				GenerateSummaryFn: func(*typgo.Context) string { return "some-summary" },
			},
			Ctx: &typgo.Context{
				Context:    createContext("-tag-name=some-tag"),
				Descriptor: &typgo.Descriptor{ProjectVersion: "9.9.9"},
			},
			RunExpectations: []*typgo.MockCommand{
				{CommandLine: "git fetch"},
			},
			Expected: &typrls.Context{
				TagName:       "some-tag",
				Alpha:         false,
				Summary:       "some-summary",
				ReleaseFolder: "release",
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			defer tt.Ctx.PatchBash(tt.RunExpectations)(t)
			var rlsCtx *typrls.Context

			tt.Releaser = typrls.NewReleaser(func(r *typrls.Context) error {
				rlsCtx = r
				return nil
			})

			err := tt.Execute(tt.Ctx)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.Expected.TagName, rlsCtx.TagName)
				require.Equal(t, tt.Expected.Alpha, rlsCtx.Alpha)
				require.Equal(t, tt.Expected.Summary, rlsCtx.Summary)
			}
		})
	}
}

func createContext(args ...string) *cli.Context {
	flagSet := flag.NewFlagSet("test", 0)
	flagSet.Bool(typrls.AlphaFlag, false, "")
	flagSet.Bool(typrls.SkipPublishFlag, false, "")
	flagSet.String(typrls.TagNameFlag, "", "")
	flagSet.String(typrls.ReleaseFolderFlag, "release", "")
	flagSet.Parse(args)
	return cli.NewContext(nil, flagSet, nil)
}

func TestReleaseCmd_Before(t *testing.T) {
	ReleaseProject := &typrls.ReleaseProject{
		Before: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}

	command := ReleaseProject.Task().CliCommand(&typgo.Descriptor{})
	require.EqualError(t, command.Before(&cli.Context{}), "some-error")
}

func TestDefaultTagFn(t *testing.T) {
	testcases := []struct {
		TestName        string
		ProjectVersion  string
		Alpha           bool
		RunExpectations []*typgo.MockCommand
		Expected        string
	}{
		{
			ProjectVersion: "0.0.1",
			Expected:       "v0.0.1",
		},
		{
			TestName:       "with alpha",
			ProjectVersion: "0.0.1",
			Alpha:          true,
			Expected:       "v0.0.1_alpha",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			c := &typgo.Context{
				Descriptor: &typgo.Descriptor{
					ProjectVersion: "0.0.1",
				},
			}
			defer c.PatchBash(tt.RunExpectations)(t)
			require.Equal(t, tt.Expected, typrls.DefaultGenerateTag(c, tt.Alpha))
		})
	}
}

func TestSummarizer(t *testing.T) {
	testCases := []struct {
		TestName        string
		RunExpectations []*typgo.MockCommand
		Expected        string
		ExpectedOut     string
	}{
		{
			TestName: "change summary",
			RunExpectations: []*typgo.MockCommand{
				{CommandLine: "git describe --tags --abbrev=0", OutputBytes: []byte("v0.0.1")},
				{CommandLine: "git --no-pager log v0.0.1..HEAD --oneline", OutputBytes: []byte("1234567 some-message-1\n1234568 some-message-3")},
			},
			Expected:    "1234567 some-message-1\n1234568 some-message-3",
			ExpectedOut: "> git describe --tags --abbrev=0\n> git --no-pager log v0.0.1..HEAD --oneline\n",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			var out strings.Builder
			c := &typgo.Context{Logger: typgo.Logger{Stdout: &out}}
			defer c.PatchBash(tt.RunExpectations)(t)
			require.Equal(t, tt.Expected, typrls.DefaultGenerateSummary(c))
			require.Equal(t, tt.ExpectedOut, out.String())
		})
	}
}

func TestChangeSummarize_HasPrefix(t *testing.T) {
	prefixes := []string{"merge", "revision"}

	testcases := []struct {
		testName string
		msg      string
		expected bool
	}{
		{msg: "Merge something", expected: true},
		{msg: "merge something", expected: true},
		{msg: "MERGE something", expected: true},
		{msg: "revision: something", expected: true},
		{msg: "asdf", expected: false},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, typrls.HasPrefix(tt.msg, prefixes))
		})
	}
}
