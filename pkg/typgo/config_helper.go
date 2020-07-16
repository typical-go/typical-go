package typgo

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/kelseyhightower/envconfig"
)

// ProcessConfig to populates the specified struct based on environment variables
func ProcessConfig(name string, spec interface{}) error {
	return envconfig.Process(name, spec)
}

func printEnv(w io.Writer, envs map[string]string) {
	color.New(color.FgGreen).Fprint(w, "ENV")
	fmt.Fprint(w, ": ")

	for key := range envs {
		fmt.Fprintf(w, "+%s ", key)
	}
	fmt.Fprintln(w)
}
