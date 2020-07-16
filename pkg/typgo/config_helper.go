package typgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/kelseyhightower/envconfig"
)

// LoadConfig to load configuration from source file
func LoadConfig(source string) (map[string]string, error) {
	file, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m := ReadConfig(file)
	for key, value := range m {
		if err := os.Setenv(key, value); err != nil {
			return nil, err
		}
	}
	return m, nil
}

// ProcessConfig to populates the specified struct based on environment variables
func ProcessConfig(name string, spec interface{}) error {
	return envconfig.Process(name, spec)
}

// ReadConfig to read config from reader
func ReadConfig(r io.Reader) (m map[string]string) {
	var ()
	m = make(map[string]string)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if i := strings.IndexRune(line, '='); i >= 0 {
			key := line[:i]
			value := line[i+1:]

			m[key] = value
		}
	}

	return
}

func printEnv(w io.Writer, envs map[string]string) {
	color.New(color.FgGreen).Fprint(w, "ENV")
	fmt.Fprint(w, ": ")

	for key := range envs {
		fmt.Fprintf(w, "+%s ", key)
	}
	fmt.Fprintln(w)
}
