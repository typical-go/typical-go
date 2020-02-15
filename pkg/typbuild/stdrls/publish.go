package stdrls

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typbuild"
)

// Publish the release
func (r *Releaser) Publish(ctx context.Context, rel *typbuild.ReleaseContext, binaries []string) (err error) {
	for _, publisher := range r.publishers {
		if err = publisher.Publish(ctx, rel, binaries); err != nil {
			return
		}
	}
	return
}
