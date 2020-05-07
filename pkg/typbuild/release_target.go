package typbuild

import (
	"errors"
	"fmt"
	"strings"
)

// ReleaseTarget contain OS and Arch as target of release operation
type ReleaseTarget string

// Validate release target
func (t ReleaseTarget) Validate() (err error) {
	s := string(t)
	if s == "" {
		return errors.New("Can't be empty")
	}
	if t.OS() == "" {
		return fmt.Errorf("Missing OS: Please make sure '%s' using 'OS/ARCH' format", s)
	}
	if t.Arch() == "" {
		return fmt.Errorf("Missing Arch: Please make sure '%s' using 'OS/ARCH' format", s)
	}
	return
}

// OS return the operating system information
func (t ReleaseTarget) OS() string {
	s := string(t)
	i := strings.Index(s, "/")
	if i < 0 {
		return ""
	}
	return s[:i]
}

// Arch return the system architecture information
func (t ReleaseTarget) Arch() string {
	s := string(t)
	i := strings.Index(s, "/")
	if i < 0 {
		return ""
	}
	return s[i+1:]
}
