package execkit

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

// PrintCommand print command
func PrintCommand(cmd *Command, w io.Writer) {
	color.New(color.FgMagenta).Fprint(w, "\n$ ")
	fmt.Fprintf(w, "%s ", cmd.Name)
	fmt.Fprintln(w, strings.Join(cmd.Args, " "))
}
