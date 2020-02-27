package typbuild

import (
	"context"
)

// Builder reponsible to build
type Builder interface {
	Build(ctx context.Context, c *Context) (bin string, err error)
}

// Prebuilder responsible to prebuild
type Prebuilder interface {
	Prebuild(ctx context.Context, c *Context) error
}
