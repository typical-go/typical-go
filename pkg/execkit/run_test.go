package execkit_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestPatch(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			Ctx:         context.Background(),
			Command:     &execkit.Command{Name: "name1", Args: []string{"arg1"}},
			ReturnError: errors.New("some-error-1"),
		},
		{
			Ctx:         context.Background(),
			Command:     &execkit.Command{Name: "name2", Args: []string{"arg2"}},
			ReturnError: errors.New("some-error-2"),
		},
	})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "name1", Args: []string{"arg1"}}),
		"some-error-1",
	)
	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "name2", Args: []string{"arg2"}}),
		"some-error-2",
	)
}

func TestPatch_CommandNoMatchedByName(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			Ctx:     context.Background(),
			Command: &execkit.Command{Name: "name1", Args: []string{"arg2"}},
		},
	})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "wrong", Args: []string{"arg2"}}),
		"execkit-mock: command not match: {wrong [arg2]} != {name1 [arg2]}",
	)
}

func TestPatch_CommandNoMatchedByArgs(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			Ctx:     context.Background(),
			Command: &execkit.Command{Name: "name2", Args: []string{"arg1", "arg2"}},
		},
	})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "name2", Args: []string{"arg2"}}),
		"execkit-mock: command not match: {name2 [arg2]} != {name2 [arg1 arg2]}",
	)
}

func TestPatch_ContextNotMatch(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			Ctx:         context.WithValue(context.Background(), "key", "vale"),
			Command:     &execkit.Command{Name: "name1", Args: []string{"arg1"}},
			ReturnError: errors.New("some-error-1"),
		},
	})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "name1", Args: []string{"arg1"}}),
		"execkit-mock: context not match",
	)
}

func TestPatch_NoRunExpectation(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "name1", Args: []string{"arg1"}}),
		"execkit-mock: no run expectation",
	)
}
