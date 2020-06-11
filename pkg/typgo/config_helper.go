package typgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// WriteConfig to write configuration to file
func WriteConfig(dest string, configs []*Configuration) (err error) {
	var fields []*Field
	for _, cfg := range configs {
		for _, field := range CreateFields(cfg) {
			fields = append(fields, field)
		}
	}

	f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	hasNewLine, err := hasLastNewLine(f)
	if err != nil {
		return
	}

	if !hasNewLine {
		fmt.Fprintln(f)
	}

	if _, err = f.Seek(0, io.SeekStart); err != nil {
		return
	}

	m := ReadConfig(f)
	for _, field := range fields {
		if _, ok := m[field.Name]; !ok {
			fmt.Fprintf(f, "%s=%v\n", field.Name, field.GetValue())
		}
	}

	return

}

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

func hasLastNewLine(f *os.File) (has bool, err error) {
	stat, err := f.Stat()
	if err != nil {
		return
	}

	if stat.Size() <= 0 {
		return true, nil
	}

	if _, err = f.Seek(-1, io.SeekEnd); err != nil {
		return
	}

	char := make([]byte, 1)
	if _, err = f.Read(char); err != nil {
		return
	}

	return (char[0] == '\n'), nil
}
