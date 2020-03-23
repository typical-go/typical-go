package typcore

import (
	"context"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore/internal/tmpl"
)

// WriteBuildToolMain to write build-tool main source
func WriteBuildToolMain(ctx context.Context, filename, descPkg string) (err error) {
	return common.WriteTemplate(
		filename,
		tmpl.BuildToolMain,
		tmpl.MainData{
			DescriptorPackage: descPkg,
		})
}

// WriteAppMain to write app main source
func WriteAppMain(ctx context.Context, filename, descPkg string) (err error) {
	return common.WriteTemplate(
		filename,
		tmpl.AppMain,
		tmpl.MainData{
			DescriptorPackage: descPkg,
		})
}
