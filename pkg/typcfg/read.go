package typcfg

import (
	"bufio"
	"io"
	"strings"
)

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
