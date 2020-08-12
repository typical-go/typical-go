package typrls

import (
	"context"
	"fmt"
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

func gitStatus(ctx context.Context) string {
	var out strings.Builder
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "git",
		Args:   []string{"status", "--porcelain"},
		Stdout: &out,
	}); err != nil {
		fmt.Fprintf(Stdout, "WARN: %s\n", err.Error())
	}
	return out.String()
}

func latestGitID(ctx context.Context) string {
	var out strings.Builder
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "git",
		Args:   []string{"rev-parse", "HEAD"},
		Stdout: &out,
	}); err != nil {
		fmt.Fprintf(Stdout, "WARN: %s\n", err.Error())
	}
	return out.String()
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
		fmt.Fprintf(Stdout, "WARN: %s\n", err.Error())
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

	var out strings.Builder
	err := execkit.Run(ctx, &execkit.Command{
		Name:   "git",
		Args:   args,
		Stdout: &out,
	})
	if err != nil {
		fmt.Fprintf(Stdout, "WARN: %s\n", err.Error())
	}

	for _, s := range strings.Split(out.String(), "\n") {
		if log := CreateLog(s); log != nil {
			logs = append(logs, log)
		}
	}
	return
}
