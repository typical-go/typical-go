package typcore

import (
	"context"

	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typcore/internal/tmpl"
)

// WriteBuildToolMain to write build-tool main source
func WriteBuildToolMain(ctx context.Context, target, descPkg string) (err error) {
	return exor.NewWriteTemplate(
		target,
		tmpl.BuildToolMain,
		tmpl.MainData{
			DescriptorPackage: descPkg,
		}).
		Execute(ctx)
}

// WriteAppMain to write app main source
func WriteAppMain(ctx context.Context, target, descPkg string) (err error) {
	return exor.NewWriteTemplate(
		target,
		tmpl.AppMain,
		tmpl.MainData{
			DescriptorPackage: descPkg,
		}).
		Execute(ctx)
}
