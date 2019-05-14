package linux

import (
	"bytes"
	"fmt"
	"os/exec"
)

// MakeDirectory execute `mkdir` in linux bash
func MakeDirectory(path string) *exec.Cmd {
	return exec.Command("mkdir", "-p", path)
}

// Download to file
func Download(url, file string) *exec.Cmd {
	return exec.Command("curl", url, "-o", file)
}

// ExtractGzip to directory
func ExtractGzip(source, dest string) *exec.Cmd {
	return exec.Command("tar", "xvzf", source, "-C", dest)
}

// Remove file
func Remove(file string) *exec.Cmd {
	return exec.Command("rm", file)
}

// Bash run bash file
func Bash(format string, v ...interface{}) *exec.Cmd {
	return exec.Command("sh", "-c", fmt.Sprintf(format, v...))
}

// Pid print process ID
func Pid(keywords ...string) (pid []byte, err error) {
	buf := bytes.Buffer{}
	buf.WriteString("ps ax")
	for _, keyword := range keywords {
		buf.WriteString(" | grep " + keyword)
	}
	buf.WriteString(" | grep -v grep")
	buf.WriteString(" | awk '{print $1}'")
	pid, err = Bash(buf.String()).Output()
	return
}
