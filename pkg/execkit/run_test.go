package execkit_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestPatch(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			CommandLine: []string{"name1", "arg1"},
			OutputBytes: []byte("some-output-bytes"),
			ErrorBytes:  []byte("some-error-bytes"),
			ReturnError: errors.New("some-error-1"),
		},
	})
	defer unpatch(t)

	var stdout strings.Builder
	var stderr strings.Builder

	require.EqualError(t,
		execkit.Run(nil, &execkit.Command{
			Name:   "name1",
			Args:   []string{"arg1"},
			Stdout: &stdout,
			Stderr: &stderr,
		}),
		"some-error-1",
	)

	require.Equal(t, "some-output-bytes", stdout.String())
	require.Equal(t, "some-error-bytes", stderr.String())
}

func TestPatch_MultipleExpectation(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"name1", "arg1"}, ReturnError: errors.New("some-error-1")},
		{CommandLine: []string{"name2", "arg2"}, ReturnError: errors.New("some-error-2")},
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
			CommandLine: []string{"name1", "arg2"},
		},
	})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "wrong", Args: []string{"arg2"}}),
		"execkit-mock: [wrong arg2] should be [name1 arg2]",
	)
}

func TestPatch_CommandNoMatchedByArgs(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			CommandLine: []string{"name2", "arg1", "arg2"},
		},
	})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "name2", Args: []string{"arg2"}}),
		"execkit-mock: [name2 arg2] should be [name2 arg1 arg2]",
	)
}

func TestPatch_NoRunExpectation(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	require.EqualError(t,
		execkit.Run(context.Background(), &execkit.Command{Name: "name1", Args: []string{"arg1"}}),
		"execkit-mock: no run expectation for [name1 arg1]",
	)
}
