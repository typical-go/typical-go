package typgo

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func writeGoSource(tmpl typtmpl.Template, target string) error {
	f, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tmpl.Execute(f); err != nil {
		return err
	}
	return goImports(target)
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
