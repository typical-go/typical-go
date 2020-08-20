package envkit

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type (
	// Map contain environment variable
	Map map[string]string
)

// Read to get environment map
func Read(r io.Reader) (m Map) {
	m = make(map[string]string)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if i := strings.IndexRune(line, '='); i >= 0 {
			m[line[:i]] = line[i+1:]
		}
	}
	return
}

// ReadFile read file to get environment map
func ReadFile(source string) (Map, error) {
	f, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Read(f), nil
}

// Save envmap to writer
func Save(m Map, w io.Writer) error {
	sortedKeys := m.Keys()
	sort.Strings(sortedKeys)
	for _, key := range sortedKeys {
		if _, err := fmt.Fprintf(w, "%s=%s\n", key, m[key]); err != nil {
			return err
		}
	}
	return nil
}

// SaveFile save envmap to file
func SaveFile(m Map, target string) error {
	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return Save(m, f)
}

// Setenv set environment variable based on map
func Setenv(m Map) error {
	for k, v := range m {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}

// Unsetenv unset environment variable
func Unsetenv(m Map) error {
	for k := range m {
		if err := os.Unsetenv(k); err != nil {
			return err
		}
	}
	return nil
}

//
// EnvMap
//

// Keys of EnvMap
func (m Map) Keys() []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
