package typrls

import (
	"context"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
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

func gitStatus(ctx context.Context, files ...string) string {
	args := []string{"status"}
	args = append(args, files...)
	args = append(args, "--porcelain")
	status, err := git(ctx, args...)
	if err != nil {
		return err.Error()
	}
	return status
}

func gitFetch(ctx context.Context) error {
	return execkit.Run(ctx, &execkit.Command{
		Name: "git",
		Args: []string{"fetch"},
	})
}

func gitTag(ctx context.Context) string {
	var out strings.Builder
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "git",
		Args:   []string{"describe", "--tags", "--abbrev=0"},
		Stdout: &out,
	}); err != nil {
		return ""
	}
	return strings.TrimSpace(out.String())
}

func gitLogs(ctx context.Context, from string) (logs []*Log) {
	var args []string
	args = append(args, "--no-pager", "log")
	if from != "" {
		args = append(args, from+"..HEAD")
	}
	args = append(args, "--oneline")

	data, err := git(ctx, args...)
	if err != nil {
		return
	}
	for _, s := range strings.Split(data, "\n") {
		if log := CreateLog(s); log != nil {
			logs = append(logs, log)
		}
	}
	return
}

// Push files to git repo
func Push(ctx context.Context, commitMessage string, files ...string) (err error) {
	args := []string{"add"}
	args = append(args, files...)
	if _, err = git(ctx, args...); err != nil {
		return
	}
	if _, err = git(ctx, "commit", "-m", commitMessage); err != nil {
		return
	}
	_, err = git(ctx, "push")
	return
}

func git(ctx context.Context, args ...string) (string, error) {
	var builder strings.Builder
	err := execkit.Run(ctx, &execkit.Command{
		Name:   "git",
		Args:   args,
		Stdout: &builder,
	})
	s := strings.TrimSuffix(builder.String(), "\n")
	return s, err
}
