package typapp

import (
	"context"
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// // WriteGoSource to write go source
// func WriteGoSource(target string, tmpl common.Template) error {
// 	if err := common.ExecuteTmplToFile(target, tmpl); err != nil {
// 		return err
// 	}
// 	goImports(target)
// 	return nil
// }

func goImports(target string) error {
	goimport := fmt.Sprintf("%s/bin/goimports", typgo.TypicalTmp)
	src := "golang.org/x/tools/cmd/goimports"

	ctx := context.Background()
	if _, err := os.Stat(goimport); os.IsNotExist(err) {
		if err := execkit.Run(ctx, &execkit.Command{
			Name:   "go",
			Args:   []string{"build", "-o", goimport, src},
			Stderr: os.Stderr,
		}); err != nil {
			return err
		}
	}

	return execkit.Run(ctx, &execkit.Command{
		Name:   goimport,
		Args:   []string{"-w", target},
		Stderr: os.Stderr,
	})
}
