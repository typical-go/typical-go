package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestBuildSysRuns(t *testing.T) {
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

	sr := typgo.BuildSysRuns{"run-1", "run-2", "run-3"}
	require.NoError(t, sr.Execute(c))
	require.Equal(t, []string{"1", "2", "3"}, seq)
}
