package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/utility/golang"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/app/common"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/urfave/cli"
)

func initiateProject(ctx *cli.Context) error {
	pkg := "github.com/typical-go/sample"
	name := filepath.Base(pkg)
	log.Infof("Remove: %s", name)
	os.RemoveAll(name)
	log.Infof("Init Project: %s", pkg)
	return runn.Execute(initproject{
		Name: name,
		Pkg:  pkg,
	})
}

type initproject struct {
	Name string
	Pkg  string
}

func (i initproject) Path(s string) string {
	return fmt.Sprintf("%s/%s", i.Name, s)
}

func (i initproject) Run() (err error) {
	return runn.Execute(
		i.appPackage,
		i.cmdPackage,
		i.dependencyPackage,
		i.typicalContext,
		i.typicalWrapper,
		i.ignoreFile,
		i.gomod,
		i.gofmt,
	)
}

func (i initproject) appPackage() error {
	return runn.Execute(
		common.Mkdir{Path: i.Path("app")},
	)
}

func (i initproject) cmdPackage() error {
	return runn.Execute(
		common.Mkdir{Path: i.Path("cmd")},
		common.Mkdir{Path: i.Path("cmd/app")},
		common.Mkdir{Path: i.Path("cmd/pre-builder")},
		common.Mkdir{Path: i.Path("cmd/build-tool")},
		common.WriteSource{Target: i.Path("cmd/app/main.go"), Source: i.appMainSrc()},
		common.WriteSource{Target: i.Path("cmd/pre-builder/main.go"), Source: i.prebuilderMainSrc()},
		common.WriteSource{Target: i.Path("cmd/build-tool/main.go"), Source: i.buildtoolMainSrc()},
	)
}

func (i initproject) appMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typapp")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Imports.Add("_", i.Pkg+"/internal/dependency")
	src.Append("typapp.Run(typical.Context)")
	return
}

func (i initproject) prebuilderMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typprebuilder")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Append("typprebuilder.Run(typical.Context)")
	return
}

func (i initproject) buildtoolMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typbuildtool")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Imports.Add("_", i.Pkg+"/internal/dependency")
	src.Append("typbuildtool.Run(typical.Context)")
	return
}

func (i initproject) typicalContext() error {
	template := `package typical

import (
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Context of Project
var Context = &typctx.Context{
	Name:    "{{.Name}}",
	Version: "0.0.1",
	Package: "{{.Pkg}}",
	Releaser: typrls.Releaser{
		Targets: []typrls.Target{"linux/amd64", "darwin/amd64"},
	},
}
`
	return runn.Execute(
		common.Mkdir{Path: i.Path("typical")},
		common.WriteTemplate{
			Target:   i.Path("typical/context.go"),
			Template: template,
			Data:     i,
		},
	)
}

func (i initproject) ignoreFile() error {
	return runn.Execute()
}

func (i initproject) dependencyPackage() error {
	return runn.Execute(
		common.Mkdir{Path: i.Path("internal/dependency")},
		common.WriteString{
			Target:     i.Path("internal/dependency/constructors.go"),
			Permission: 0644,
			Content:    "package dependency",
		},
	)
}

func (i initproject) typicalWrapper() error {
	content := `#!/bin/bash
set -e

BIN=${TYPICAL_BIN:-bin}
CMD=${TYPICAL_CMD:-cmd}
BUILD_TOOL=${TYPICAL_BUILD_TOOL:-build-tool}
PRE_BUILDER=${TYPICAL_PRE_BUILDER:-pre-builder}
METADATA=${TYPICAL_METADATA:-.typical-metadata}

CHECKSUM_PATH="$METADATA/checksum "
CHECKSUM_DATA=$(cksum typical/context.go)

if ! [ -s .typical-metadata/checksum ]; then
	mkdir -p $METADATA
	cksum typical/context.go > $CHECKSUM_PATH
else
	CHECKSUM_UPDATED=$([ "$CHECKSUM_DATA" == "$(cat $CHECKSUM_PATH )" ] ; echo $?)
fi

if [ "$CHECKSUM_UPDATED" == "1" ] || ! [[ -f $BIN/$PRE_BUILDER ]] ; then 
	echo $CHECKSUM_DATA > $CHECKSUM_PATH
	echo "Build the pre-builder"
	go build -o $BIN/$PRE_BUILDER ./$CMD/$PRE_BUILDER
fi

./$BIN/$PRE_BUILDER $CHECKSUM_UPDATED
./$BIN/$BUILD_TOOL $@`
	return runn.Execute(
		common.WriteString{
			Target:     i.Path("typicalw"),
			Permission: 0700,
			Content:    content,
		},
	)
}

func (i initproject) gomod() (err error) {
	template := `module {{.Pkg}}

go 1.12

require github.com/typical-go/typical-go v{{.TypicalVersion}}
`
	return runn.Execute(
		common.WriteTemplate{
			Target:   i.Path("go.mod"),
			Template: template,
			Data: struct {
				Pkg            string
				TypicalVersion string
			}{
				Pkg:            i.Pkg,
				TypicalVersion: Version,
			},
		},
	)
}

func (i initproject) gofmt() (err error) {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = i.Name
	return cmd.Run()
}
