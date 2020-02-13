package stdrelease

import (
	"context"
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
)

// Tag return relase tag
func (r *Releaser) Tag(ctx context.Context, version string, alpha bool) string {
	var b strings.Builder
	b.WriteString("v")
	b.WriteString(version)
	if r.IncludeBranch {
		b.WriteString("_")
		b.WriteString(git.Branch(ctx))
	}
	if r.IncludeCommitID {
		b.WriteString("_")
		b.WriteString(git.LatestCommit(ctx))
	}
	if alpha {
		b.WriteString("_alpha")
	}
	return b.String()
}
