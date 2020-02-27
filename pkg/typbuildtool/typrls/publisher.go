package typrls

import (
	"context"

	"github.com/typical-go/typical-go/pkg/git"
)

// Publisher reponsible to publish the release to external source
type Publisher interface {
	Publish(context.Context, *PublishContext) error
}

// PublishContext is context of publish
type PublishContext struct {
	*Context
	Tag      string
	Binaries []string
	GitLogs  []*git.Log
}
