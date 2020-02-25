package typicalgo

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/runnerkit"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo/internal/tmpl"
)

func constructProject(ctx context.Context, pkg string) (err error) {
	name := filepath.Base(pkg)
	if common.IsFileExist(name) {
		return fmt.Errorf("'%s' already exist", name)
	}
	return runnerkit.Run(ctx,
		constructproj{
			TemplateData: tmpl.TemplateData{
				Name: name,
				Pkg:  pkg,
			},
			ctx: ctx,
		},
	)
}

type constructproj struct {
	tmpl.TemplateData
	ctx context.Context
}

func (i constructproj) Path(s string) string {
	return fmt.Sprintf("%s/%s", i.Name, s)
}

func (i constructproj) Run(ctx context.Context) (err error) {
	return runnerkit.Run(ctx,
		i.appPackage,
		i.cmdPackage,
		i.descriptor,
		i.ignoreFile,
		wrapper(i.Name, i.Pkg),
		runnerkit.GoFmt("./..."),
		i.gomod,
	)
}

func (i constructproj) appPackage(ctx context.Context) error {
	return runnerkit.Run(ctx, runnerkit.Mkdir(i.Path("app")))
}

func (i constructproj) descriptor(ctx context.Context) error {
	return runnerkit.Run(ctx,
		runnerkit.Mkdir(i.Path("typical")),
		runnerkit.WriteTemplate(i.Path("typical/descriptor.go"), tmpl.Descriptor, i.TemplateData, 0666),
	)
}

func (i constructproj) cmdPackage(ctx context.Context) error {
	appMainPath := fmt.Sprintf("%s/%s", typcore.DefaultCmdFolder, i.Name)
	data := tmpl.MainSrcData{
		DescriptorPackage: i.Pkg + "/typical",
	}
	return runnerkit.Run(ctx,
		runnerkit.Mkdir(i.Path(typcore.DefaultCmdFolder)),
		runnerkit.Mkdir(i.Path(appMainPath)),
		runnerkit.WriteTemplate(i.Path(appMainPath+"/main.go"), tmpl.MainSrcApp, data, 0666),
	)
}

func (i constructproj) ignoreFile(ctx context.Context) error {
	return runnerkit.Run(ctx,
		runnerkit.WriteString(i.Path(".gitignore"), tmpl.Gitignore, 0700),
	)
}

func (i constructproj) gomod(ctx context.Context) (err error) {
	return runnerkit.Run(ctx,
		runnerkit.WriteTemplate(i.Path("go.mod"), tmpl.GoMod, tmpl.GoModData{
			Pkg:            i.Pkg,
			TypicalVersion: typcore.Version,
		}, 0666),
	)
}

func wrapper(path, pkg string) runnerkit.Runner {
	return runnerkit.WriteTemplate(
		path+"/typicalw",
		tmpl.Typicalw,
		tmpl.TypicalwData{
			DescriptorPackage: fmt.Sprintf("%s/typical", pkg),
			DescriptorFile:    "typical/descriptor.go",
			ChecksumFile:      typcore.DefaultTempFolder + "/checksum",
			LayoutTemp:        typcore.DefaultTempFolder,
		},
		0700,
	)
}
