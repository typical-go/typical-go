package typrls

import (
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// Git detail
	Git struct {
		Status     string
		CurrentTag string
		Logs       []*Log
	}
	// Log git
	Log struct {
		ShortCode    string
		Message      string
		CoAuthoredBy string
	}
)

// CreateLog to create git log from raw message
func CreateLog(raw string) *Log {
	if len(raw) < 7 {
		return nil
	}
	raw = strings.TrimSpace(raw)
	i := strings.Index(raw, " ")
	if i != 7 {
		return nil
	}
	message := raw[7:]
	coAuthoredBy := ""
	if i := strings.Index(message, "Co-Authored-By:"); i >= 0 {
		coAuthoredBy = message[i+len("Co-Authored-By:"):]
		message = message[:i]

	}
	return &Log{
		ShortCode:    strings.TrimSpace(raw[:7]),
		Message:      strings.TrimSpace(message),
		CoAuthoredBy: strings.TrimSpace(coAuthoredBy),
	}
}

// GitFetch fetch latest update
func GitFetch(c *typgo.Context) error {
	return c.Execute(&typgo.Bash{Name: "git", Args: []string{"fetch"}, Stderr: os.Stderr})
}

// GitCurrentTag return git current tag
func GitCurrentTag(c *typgo.Context) string {
	var out strings.Builder
	if err := c.Execute(&typgo.Bash{
		Name:   "git",
		Args:   []string{"describe", "--tags", "--abbrev=0"},
		Stdout: &out,
	}); err != nil {
		c.Infof("WARN: %s\n", err.Error())
	}
	return strings.TrimSpace(out.String())
}

// GitLogs return git logs
func GitLogs(c *typgo.Context, from string) (logs []*Log) {
	var args []string
	args = append(args, "--no-pager", "log")
	if from != "" {
		args = append(args, from+"..HEAD")
	}
	args = append(args, "--oneline")

	var out strings.Builder
	err := c.Execute(&typgo.Bash{Name: "git", Args: args, Stdout: &out})
	if err != nil {
		c.Infof("WARN: %s\n", err.Error())
	}

	for _, s := range strings.Split(out.String(), "\n") {
		if log := CreateLog(s); log != nil {
			logs = append(logs, log)
		}
	}
	return
}
