package typgo

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func writeGoSource(target string, tmpl typtmpl.Template) error {
	if err := typtmpl.ExecuteToFile(target, tmpl); err != nil {
		return err
	}
	goImports(target)
	return nil
}

func goImports(target string) error {
	goimport := fmt.Sprintf("%s/bin/goimports", TypicalTmp)
	src := "golang.org/x/tools/cmd/goimports"

	if _, err := os.Stat(goimport); os.IsNotExist(err) {
		cmd := exec.Command("go", "build", "-o", goimport, src)
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	cmd := exec.Command(goimport, "-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
