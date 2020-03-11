package git

import (
	"context"
	"os/exec"
	"strings"
)

// Status is same with `git status --porcelain`
func Status(ctx context.Context, files ...string) string {
	args := []string{"status"}
	args = append(args, files...)
	args = append(args, "--porcelain")
	status, err := git(ctx, args...)
	if err != nil {
		return err.Error()
	}
	return status
}

// Fetch is same with `git fetch`
func Fetch(ctx context.Context) error {
	return exec.CommandContext(ctx, "git", "fetch").Run()
}

// LatestTag to get latest tag and its hash key
func LatestTag(ctx context.Context) string {
	tag, err := git(ctx, "describe", "--tags", "--abbrev=0")
	if err != nil {
		return ""
	}
	return tag
}

// RetrieveLogs to get git logs
func RetrieveLogs(ctx context.Context, from string) (logs []*Log) {
	var (
		data string
		err  error
		args []string
	)

	args = append(args, "--no-pager", "log")
	if from != "" {
		args = append(args, from+"..HEAD")
	}
	args = append(args, "--oneline")

	if data, err = git(ctx, args...); err != nil {
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

// Branch to return current branch
func Branch(ctx context.Context) string {
	var (
		branch string
		err    error
	)
	if branch, err = git(ctx, "rev-parse", "--abbrev-ref", "HEAD"); err != nil {
		return ""
	}
	return branch
}

// LatestCommit return latest commit in short hash
func LatestCommit(ctx context.Context) string {
	var (
		commit string
		err    error
	)
	if commit, err = git(ctx, "rev-parse", "--short", "HEAD"); err != nil {
		return ""
	}
	return commit
}

func git(ctx context.Context, args ...string) (s string, err error) {
	var builder strings.Builder
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Stdout = &builder
	cmd.Stderr = &builder
	err = cmd.Run()
	s = strings.TrimSuffix(builder.String(), "\n")
	return
}
