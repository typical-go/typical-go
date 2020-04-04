package typcfg

import (
	"bufio"
	"os"
	"strings"
)

// Load configuration from source file
func Load(source string) (m map[string]string, err error) {
	var (
		b strings.Builder
	)
	if m, err = Read(source); err != nil {
		return
	}

	if len(m) > 0 {
		for key, value := range m {
			if err = os.Setenv(key, value); err != nil {
				return
			}
			b.WriteString("+" + key + " ")
		}
	}

	return
}

// Read config file
func Read(source string) (m map[string]string, err error) {
	var (
		file *os.File
	)
	m = make(map[string]string)

	if file, err = os.Open(source); err != nil {
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		i := strings.IndexRune(line, '=')
		key := line[:i]
		value := line[i+1:]

		m[key] = value
	}

	return
}
