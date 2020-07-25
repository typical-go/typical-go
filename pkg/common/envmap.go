package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
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
func Setenv(m EnvMap) {
	w := os.Stdout
	color.New(color.FgGreen).Fprint(w, "ENV")
	fmt.Fprint(w, ": ")
	defer fmt.Fprintln(w)
	for k, v := range m {
		if err := os.Setenv(k, v); err != nil {
			fmt.Fprintf(w, "failed: %s ", err.Error())
			return
		}
		fmt.Fprintf(w, "+%s ", k)
	}
}

// Unsetenv unset environment variable
func Unsetenv(m EnvMap) error {
	for k := range m {
		if err := os.Unsetenv(k); err != nil {
			return err
		}
	}
	return nil
}

// LoadEnv to setenv from file
func LoadEnv(filename string) error {
	envmap, err := CreateEnvMapFromFile(filename)
	if err != nil {
		return err
	}
	Setenv(envmap)
	return nil
}

//
// EnvMap
//

// Keys of EnvMap
func (m EnvMap) Keys() []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Save envmap to writer
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

// SaveToFile save envmap to file
func (m EnvMap) SaveToFile(target string) error {
	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return m.Save(f)
}
