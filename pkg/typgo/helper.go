package typgo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typvar"
)

func execute(c *Context, cmd *execkit.Command) error {
	cmd.Print(os.Stdout)
	fmt.Fprintln(os.Stdout)

	return cmd.Run(c.Ctx())
}

// ExcludeMessage return true is message mean to be exclude
func ExcludeMessage(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range typvar.ExclMsgPrefix {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}

// WalkLayout return dirs and files
func WalkLayout(layouts []string) (dirs, files []string) {
	for _, layout := range layouts {
		filepath.Walk(layout, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}

			if info.IsDir() {
				dirs = append(dirs, path)
				return nil
			}

			if isGoSource(path) {
				files = append(files, path)
			}
			return nil
		})
	}
	return
}

func isGoSource(path string) bool {
	return strings.HasSuffix(path, ".go") &&
		!strings.HasSuffix(path, "_test.go")
}
