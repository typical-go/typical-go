package typgo

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/typical-go/typical-go/pkg/oskit"
)

type (
	// TaskNames run command in current BuildSys
	TaskNames []string
)

var _ Action = (TaskNames)(nil)

// Execute BuildCmdRuns
func (r TaskNames) Execute(c *Context) error {
	for _, name := range r {
		taskSignature(oskit.Stdout, fmt.Sprintf("./typicalw %s", name))
		if err := RunBash(c.Ctx(), &Bash{
			Name:   "./typicalw",
			Args:   strings.Split(name, " "),
			Stdout: oskit.Stdout,
			Stderr: os.Stderr,
			Stdin:  os.Stdin,
		}); err != nil {
			return err
		}
	}
	return nil
}

func taskSignature(w io.Writer, task string) {
	fmt.Fprintf(w, "\n----- %s: ", ProjectName)
	color.New(color.FgCyan).Fprint(w, task)
	fmt.Fprint(w, " -----\n\n")
}
