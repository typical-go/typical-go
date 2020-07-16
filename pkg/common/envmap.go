package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type (
	// EnvMap map contain environment variable
	EnvMap map[string]string
)

// CreateEnvMap create EnvMap instance from reader
func CreateEnvMap(r io.Reader) (m EnvMap) {
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

// CreateEnvMapFromFile to create EnvMap from file
func CreateEnvMapFromFile(source string) (EnvMap, error) {
	f, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return CreateEnvMap(f), nil
}

// Setenv set environment variable based on map
func (m EnvMap) Setenv() error {
	for k, v := range m {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}

// Unsetenv unset environment variable
func (m EnvMap) Unsetenv() error {
	for k := range m {
		if err := os.Unsetenv(k); err != nil {
			return err
		}
	}
	return nil
}

// Keys of EnvMap
func (m EnvMap) Keys() []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Save map to writer
func (m EnvMap) Save(w io.Writer) error {
	sortedKeys := m.Keys()
	sort.Strings(sortedKeys)
	for _, key := range sortedKeys {
		if _, err := fmt.Fprintf(w, "%s=%s\n", key, m[key]); err != nil {
			return err
		}
	}
	return nil
}
