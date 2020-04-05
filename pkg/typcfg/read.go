package typcfg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Load configuration from source file
func Load(source string) (m map[string]string, err error) {
	var (
		file *os.File
	)

	if file, err = os.Open(source); err != nil {
		return
	}
	defer file.Close()

	m = Read(file)
	for key, value := range m {
		if err = os.Setenv(key, value); err != nil {
			return
		}
	}

	return
}

// Read config file
func Read(r io.Reader) (m map[string]string) {
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
