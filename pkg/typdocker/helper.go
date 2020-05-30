package typdocker

import (
	"context"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
)

// Major version from docker-composer version
func Major(version string) string {
	i := strings.IndexRune(version, '.')
	if i < 0 {
		return version
	}

	return version[:i]
}

func dockerIDs(ctx context.Context) (ids []string, err error) {
	var out strings.Builder
	cmd := &execkit.Command{
		Name:   "docker",
		Args:   []string{"ps", "-q"},
		Stderr: os.Stderr,
		Stdout: &out,
	}

	execkit.PrintCommand(cmd, os.Stdout)

	if err = cmd.Run(ctx); err != nil {
		return
	}

	for _, id := range strings.Split(out.String(), "\n") {
		if id != "" {
			ids = append(ids, id)
		}
	}
	return
}

func kill(ctx context.Context, id string) (err error) {
	cmd := &execkit.Command{
		Name:   "docker",
		Args:   []string{"kill", id},
		Stderr: os.Stderr,
	}
	execkit.PrintCommand(cmd, os.Stdout)
	return cmd.Run(ctx)
}
