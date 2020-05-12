package typcfg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// WriteConfig to write configuration to file
func WriteConfig(dest string, c Configurer) (err error) {
	var (
		fields []*Field
		f      *os.File
		m      map[string]string
	)

	for _, cfg := range c.Configurations() {
		for _, field := range CreateFields(cfg) {
			fields = append(fields, field)
		}
	}

	if f, err = os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err != nil {
		return
	}
	defer f.Close()

	stat, _ := f.Stat()
	if stat.Size() > 0 {
		fmt.Fprintln(f)
	}

	m = ReadConfig(f)
	for _, field := range fields {
		if _, ok := m[field.Name]; !ok {
			fmt.Fprintf(f, "%s=%v\n", field.Name, field.GetValue())
		}
	}

	return

}

// LoadConfig to load configuration from source file
func LoadConfig(source string) (m map[string]string, err error) {
	var file *os.File

	if file, err = os.Open(source); err != nil {
		return
	}
	defer file.Close()

	for key, value := range ReadConfig(file) {
		if err = os.Setenv(key, value); err != nil {
			return
		}
	}
	return
}

// // ProcessConfig to populates the specified struct based on environment variables
// func ProcessConfig(name string, spec interface{}) error {
// 	return envconfig.Process(name, spec)
// }

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
