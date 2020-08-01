package typgo

import (
	"context"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
)

// GoImports run goimports
func GoImports(target string) error {
	ctx := context.Background()
	goimports, err := InstallTool(ctx, "goimports", "golang.org/x/tools/cmd/goimports")
	if err != nil {
		return err
	}
	return execkit.Run(ctx, &execkit.Command{
		Name:   goimports,
		Args:   []string{"-w", target},
		Stderr: os.Stderr,
	})
}

// InstallTool install tool to typical-tmp folder
func InstallTool(ctx context.Context, name, source string) (string, error) {
	output := fmt.Sprintf("%s/bin/%s", TypicalTmp, name)
	if _, err := os.Stat(output); os.IsNotExist(err) {
		gobuild := &execkit.GoBuild{
			Output:      output,
			MainPackage: source,
		}
		if err := execkit.Run(ctx, gobuild); err != nil {
			return "", err
		}
	}
	return output, nil
}
