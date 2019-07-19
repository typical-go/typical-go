package command

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/typical-go/typical-go/typicore"
)

func execute(cmds ...typicore.Statement) (err error) {
	for _, cmd := range cmds {
		execCmd, ok := cmd.(*exec.Cmd)
		if ok {
			buf := bytes.Buffer{}
			execCmd.Stdout = &buf
			execCmd.Stderr = &buf

			err := execCmd.Run()
			if err != nil {
				return fmt.Errorf("%s: %s", err.Error(), buf.String())
			}
		} else {
			err = cmd.Run()
		}

		if err != nil {
			return
		}
	}
	return
}
