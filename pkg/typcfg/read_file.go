package typcfg

import (
	"bufio"
	"os"
	"strings"
)

// ReadFile to read config file
func ReadFile(source string) (m map[string]string, err error) {
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
