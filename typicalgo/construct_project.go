package typicalgo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo/internal/tmpl"
)

func constructProject(ctx context.Context, pkg string) (err error) {
	name := filepath.Base(pkg)

	if _, err = os.Stat(name); os.IsNotExist(err) {
		return
	}

	return exor.Execute(ctx,
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

func (i constructproj) Execute(ctx context.Context) (err error) {
	return exor.Execute(ctx,
		exor.New(i.appPackage),
		exor.New(i.cmdPackage),
		exor.New(i.descriptor),
		exor.New(i.ignoreFile),
		wrapper(i.Name, i.Pkg),
		exor.NewGoFmt("./..."),
		exor.New(i.gomod),
	)
}

func (i constructproj) appPackage(ctx context.Context) error {
	return exor.Execute(ctx, exor.NewMkdir(i.Path("app")))
}

func (i constructproj) descriptor(ctx context.Context) error {
	return exor.Execute(ctx,
		exor.NewMkdir(i.Path("typical")),
		exor.NewWriteTemplate(i.Path("typical/descriptor.go"), tmpl.Descriptor, i.TemplateData),
	)
}

func (i constructproj) cmdPackage(ctx context.Context) error {
	appMainPath := fmt.Sprintf("%s/%s", typcore.DefaultCmdFolder, i.Name)
	// data := tmpl.MainSrcData{
	// 	DescriptorPackage: i.Pkg + "/typical",
	// }
	return exor.Execute(ctx,
		exor.NewMkdir(i.Path(typcore.DefaultCmdFolder)),
		exor.NewMkdir(i.Path(appMainPath)),
		// exor.NewWriteTemplate(i.Path(appMainPath+"/main.go"), tmpl.MainSrcApp, data),
	)
}

func (i constructproj) ignoreFile(ctx context.Context) error {
	return exor.Execute(ctx,
		exor.NewWriteString(i.Path(".gitignore"), tmpl.Gitignore),
	)
}

func (i constructproj) gomod(ctx context.Context) (err error) {
	return exor.Execute(ctx,
		exor.NewWriteTemplate(i.Path("go.mod"), tmpl.GoMod, tmpl.GoModData{
			Pkg:            i.Pkg,
			TypicalVersion: typcore.Version,
		}),
	)
}

func wrapper(path, pkg string) exor.Executor {
	return exor.NewWriteTemplate(
		path+"/typicalw",
		tmpl.Typicalw,
		tmpl.TypicalwData{
			DescriptorPackage: fmt.Sprintf("%s/typical", pkg),
			DescriptorFile:    "typical/descriptor.go",
			ChecksumFile:      typcore.DefaultTempFolder + "/checksum",
			LayoutTemp:        typcore.DefaultTempFolder,
		},
	)
}
