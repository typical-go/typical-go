package typgo

import (
	"fmt"
	"os"
)

// GoImports run goimports
func GoImports(c *Context, target string) error {

	goimports, err := InstallTool(c, "goimports", "golang.org/x/tools/cmd/goimports")
	if err != nil {
		return err
	}
	return c.Execute(&Bash{
		Name:   goimports,
		Args:   []string{"-w", target},
		Stderr: os.Stderr,
	})
}

// InstallTool install tool to typical-tmp folder
func InstallTool(c *Context, name, source string) (string, error) {
	output := fmt.Sprintf("%s/bin/%s", TypicalTmp, name)

	if _, err := os.Stat(output); os.IsNotExist(err) {
		if err := c.Execute(&GoBuild{
			Output:      output,
			MainPackage: source,
		}); err != nil {
			return "", err
		}
	}
	return output, nil
}
