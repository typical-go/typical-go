package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestBuildCmdRuns(t *testing.T) {
	var seq []string
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{
			Commands: []*cli.Command{
				{Name: "run-1", Action: func(*cli.Context) error { seq = append(seq, "1"); return nil }},
				{Name: "run-2", Action: func(*cli.Context) error { seq = append(seq, "2"); return nil }},
				{Name: "run-3", Action: func(*cli.Context) error { seq = append(seq, "3"); return nil }},
			},
		},
	}

	sr := typgo.BuildCmdRuns{"run-1", "run-2", "run-3"}
	require.NoError(t, sr.Execute(c))
	require.Equal(t, []string{"1", "2", "3"}, seq)
}

func TestBuildCmdRuns_Error(t *testing.T) {
	var seq []string
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{
			Commands: []*cli.Command{
				{Name: "run-1", Action: func(*cli.Context) error { seq = append(seq, "1"); return errors.New("some-error") }},
				{Name: "run-2", Action: func(*cli.Context) error { seq = append(seq, "2"); return nil }},
				{Name: "run-3", Action: func(*cli.Context) error { seq = append(seq, "3"); return nil }},
			},
		},
	}

	sr := typgo.BuildCmdRuns{"run-1", "run-2", "run-3"}
	require.EqualError(t, sr.Execute(c), "some-error")
	require.Equal(t, []string{"1"}, seq)
}

func TestRunCommand(t *testing.T) {
	c := &typgo.Context{
		Context: &cli.Context{},
		BuildSys: &typgo.BuildSys{
			Commands: []*cli.Command{
				{
					Name: "cmd1",
					Subcommands: []*cli.Command{
						{
							Name:   "sub1",
							Action: func(*cli.Context) error { return errors.New("sub1-error") },
						},
						{
							Name: "sub2",
							Subcommands: []*cli.Command{
								{
									Name:   "ssub1",
									Action: func(*cli.Context) error { return errors.New("ssub1-error") },
								},
							},
							Action: func(*cli.Context) error { return errors.New("sub2-error") },
						},
					},
					Action: func(*cli.Context) error { return errors.New("cmd1-error") },
				},
				{
					Name:   "cmd2",
					Action: func(*cli.Context) error { return errors.New("cmd2-error") },
				},
				{
					Name:   "cmd3",
					Action: func(*cli.Context) error { return errors.New("cmd3-error") },
				},
			},
		},
	}
	require.EqualError(t, typgo.RunCommand(c, "noname"), "typgo: noname not found")
	require.EqualError(t, typgo.RunCommand(c, "cmd1"), "cmd1-error")
	require.EqualError(t, typgo.RunCommand(c, "cmd2"), "cmd2-error")
	require.EqualError(t, typgo.RunCommand(c, "cmd3"), "cmd3-error")
	require.EqualError(t, typgo.RunCommand(c, "cmd1.sub1"), "sub1-error")
	require.EqualError(t, typgo.RunCommand(c, "cmd1.sub2"), "sub2-error")
	require.EqualError(t, typgo.RunCommand(c, "cmd1.sub2.ssub1"), "ssub1-error")
	require.EqualError(t, typgo.RunCommand(c, "cmd1.sub2.ssub2"), "typgo: cmd1.sub2.ssub2 not found")

}
