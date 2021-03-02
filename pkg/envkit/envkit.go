package envkit

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Read to get environment map
func Read(r io.Reader) map[string]string {
	m := make(map[string]string)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if i := strings.IndexRune(line, '='); i >= 0 {
			m[line[:i]] = line[i+1:]
		}
	}
	return m
}

// ReadFile read file to get environment map
func ReadFile(source string) (map[string]string, error) {
	f, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Read(f), nil
}

// Save envmap to writer
func Save(m map[string]string, w io.Writer) error {
	for _, key := range SortedKeys(m) {
		if _, err := fmt.Fprintf(w, "%s=%s\n", key, m[key]); err != nil {
			return err
		}
	}
	return nil
}

// SaveFile save envmap to file
func SaveFile(m map[string]string, target string) error {
	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return Save(m, f)
}

// Setenv set environment variable based on map
func Setenv(m map[string]string) error {
	for k, v := range m {
		if v != "" {
			if err := os.Setenv(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Unsetenv unset environment variable
func Unsetenv(m map[string]string) error {
	for k := range m {
		if err := os.Unsetenv(k); err != nil {
			return err
		}
	}
	return nil
}

// SortedKeys of EnvMap
func SortedKeys(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
