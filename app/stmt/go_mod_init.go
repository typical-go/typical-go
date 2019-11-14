package stmt

import (
	"fmt"
	"go/build"
	"os/exec"
)

type GoModInit struct {
	ProjectPath string
	PackageName string
}

func (i GoModInit) Run() error {
	goCommand := fmt.Sprintf("%s/bin/go", build.Default.GOROOT)
	cmd := exec.Command(goCommand, "mod", "init", i.PackageName)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stdout
	return cmd.Run()
}
