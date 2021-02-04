package typgo_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestTaskNames(t *testing.T) {
	var out strings.Builder

	app := typgo.BuildTool(&typgo.Descriptor{
		Tasks: []typgo.Tasker{
			&typgo.Task{
				Name: "task-1",
				Action: typgo.NewAction(func(*typgo.Context) error {
					fmt.Fprintln(&out, "a")
					return nil
				}),
			},
			&typgo.Task{
				Name: "task-2",
				Action: typgo.NewAction(func(*typgo.Context) error {
					fmt.Fprintln(&out, "b")
					return nil
				}),
			},
			&typgo.Task{
				Name:   "all",
				Action: typgo.TaskNames{"task-1", "task-2"},
			},
		},
	})

	require.NoError(t, app.Run([]string{"typical-build", "all"}))
	require.Equal(t, "a\nb\n", out.String())

}

func TestTaskNames_FistTaskError(t *testing.T) {
	var out strings.Builder

	app := typgo.BuildTool(&typgo.Descriptor{
		Tasks: []typgo.Tasker{
			&typgo.Task{
				Name: "task-1",
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			&typgo.Task{
				Name: "task-2",
				Action: typgo.NewAction(func(*typgo.Context) error {
					fmt.Fprintln(&out, "b")
					return nil
				}),
			},
			&typgo.Task{
				Name:   "all",
				Action: typgo.TaskNames{"task-1", "task-2"},
			},
		},
	})

	require.EqualError(t, app.Run([]string{"typical-build", "all"}), "some-error")
}
