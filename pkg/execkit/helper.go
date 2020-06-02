package execkit

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

// PrintCommand print command
func PrintCommand(cmd *Command, w io.Writer) {
	color.New(color.FgMagenta).Fprint(w, "\n$ ")
	fmt.Fprintln(w, cmd)
}
