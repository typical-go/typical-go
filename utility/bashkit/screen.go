package bashkit

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Screen immitate bash screen
type Screen interface {
	Run(cmds ...*exec.Cmd) error
}

// NewScreen return new instance of screen
func NewScreen() Screen {
	return &screen{}
}

type screen struct {
}

func (s *screen) Run(cmds ...*exec.Cmd) (err error) {
	for _, cmd := range cmds {
		buf := bytes.Buffer{}
		cmd.Stdout = &buf
		cmd.Stderr = &buf

		// change writer
		fmt.Println(strings.Join(cmd.Args, " "))

		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("%s: %s\n", err.Error(), buf.String())
		}
	}
	return nil
}
