package typbuildtool

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"
)

func prebuild(ctx context.Context, d *typcore.ProjectDescriptor) (err error) {
	var (
		stdPrebuilder typcore.StandardPrebuilder
		pc            *typcore.PrebuildContext
	)
	if pc, err = typcore.CreatePrebuildContext(ctx, d); err != nil {
		return
	}
	if err = stdPrebuilder.Prebuild(pc); err != nil {
		return
	}
	return
}
