package typgo

import (
	"fmt"
	"io"
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
		args := []string{c.App.Name}
		args = append(args, strings.Split(name, " ")...)

		if err := c.App.Run(args); err != nil {
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
