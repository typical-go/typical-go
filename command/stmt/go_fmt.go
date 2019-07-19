package stmt

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
)

type GoFmt struct {
	ProjectPath string
}

func (i GoFmt) Run() error {
	os.Chdir(i.ProjectPath)
	goCommand := fmt.Sprintf("%s/bin/go", build.Default.GOROOT)
	cmd := exec.Command(goCommand, "fmt", "./...")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stdout
	return cmd.Run()
}
