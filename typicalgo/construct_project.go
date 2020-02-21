package typicalgo

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/common/stdrun"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo/internal/tmpl"
)

func constructProject(ctx context.Context, pkg string) (err error) {
	name := filepath.Base(pkg)
	if common.IsFileExist(name) {
		return fmt.Errorf("'%s' already exist", name)
	}
	return common.Run(constructproj{
		TemplateData: tmpl.TemplateData{
			Name: name,
			Pkg:  pkg,
		},
		ctx: ctx,
	})
}

type constructproj struct {
	tmpl.TemplateData
	ctx context.Context
}

func (i constructproj) Path(s string) string {
	return fmt.Sprintf("%s/%s", i.Name, s)
}

func (i constructproj) Run() (err error) {
	return common.Run(
		i.appPackage,
		i.cmdPackage,
		i.descriptor,
		i.ignoreFile,
		wrapper(i.Name, i.Pkg),
		stdrun.NewGoFmt(i.ctx, "./..."),
		i.gomod,
	)
}

func (i constructproj) appPackage() error {
	stmts := []interface{}{
		stdrun.NewMkdir(i.Path("app")),
	}

	return common.Run(stmts...)
}

func (i constructproj) descriptor() error {
	return common.Run(
		stdrun.NewMkdir(i.Path("typical")),
		stdrun.NewWriteTemplate(i.Path("typical/descriptor.go"), tmpl.Descriptor, i.TemplateData),
	)
}

func (i constructproj) cmdPackage() error {
	appMainPath := fmt.Sprintf("%s/%s", typcore.DefaultLayout.Cmd, i.Name)
	data := tmpl.MainSrcData{
		DescriptorPackage: i.Pkg + "/typical",
	}
	return common.Run(
		stdrun.NewMkdir(i.Path(typcore.DefaultLayout.Cmd)),
		stdrun.NewMkdir(i.Path(appMainPath)),
		stdrun.NewWriteTemplate(i.Path(appMainPath+"/main.go"), tmpl.MainSrcApp, data),
	)
}

func (i constructproj) ignoreFile() error {
	return common.Run(
		stdrun.NewWriteString(i.Path(".gitignore"), tmpl.Gitignore).WithPermission(0700),
	)
}

func (i constructproj) gomod() (err error) {
	return common.Run(
		stdrun.NewWriteTemplate(i.Path("go.mod"), tmpl.GoMod, tmpl.GoModData{
			Pkg:            i.Pkg,
			TypicalVersion: Version,
		}),
	)
}

func wrapper(path, pkg string) common.Runner {
	return stdrun.NewWriteTemplate(
		path+"/typicalw",
		tmpl.Typicalw,
		tmpl.TypicalwData{
			DescriptorPackage: fmt.Sprintf("%s/typical", pkg),
			DescriptorFile:    "typical/descriptor.go",
			ChecksumFile:      ".typical-tmp/checksum",
			LayoutTemp:        typcore.DefaultLayout.Temp,
		},
	).WithPermission(0700)
}
