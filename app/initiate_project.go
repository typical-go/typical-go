package app

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/app/common"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/urfave/cli"
)

func initiateProject(ctx *cli.Context) error {
	parent := "sample"
	pkg := "github.com/typical-go/hello-world"
	log.Infof("Remove: %s", parent)
	os.RemoveAll(parent)
	log.Infof("Init Project: %s", pkg)
	return runn.Execute(initproject{
		Name:   name(pkg),
		Pkg:    pkg,
		Parent: parent,
	})
}

func name(pkg string) string {
	// TODO: handle window path format
	chunks := strings.Split(pkg, "/")
	return chunks[len(chunks)-1]
}

type initproject struct {
	Name   string
	Parent string
	Pkg    string
}

func (i initproject) Path(s string) string {
	return fmt.Sprintf("%s/%s/%s", i.Parent, i.Name, s)
}

func (i initproject) Run() (err error) {
	return runn.Execute(
		i.generateAppPackage,
		i.generateCmdPackage,
		i.generateTypicalContext,
		i.generateIgnoreFile,
		i.generateTypicalWrapper,
	)
}

func (i initproject) generateAppPackage() error {
	log.Info("Generate App Package")
	return runn.Execute(
		common.Mkdir{Path: i.Path("app")},
	)
}

func (i initproject) generateCmdPackage() error {
	log.Info("Generate Cmd Package")
	// 	_ "github.com/typical-go/typical-go/internal/dependency"
	// 	"github.com/typical-go/typical-go/pkg/typapp"
	// 	"github.com/typical-go/typical-go/typical"

	return runn.Execute(
		common.Mkdir{Path: i.Path("cmd")},
		common.Mkdir{Path: i.Path("cmd/app")},
		common.Mkdir{Path: i.Path("cmd/pre-builder")},
		common.Mkdir{Path: i.Path("cmd/build-tool")},
	)
}

func (i initproject) generateTypicalContext() error {
	log.Info("Generate Typical Context")
	return runn.Execute(
		common.Mkdir{Path: i.Path("typical")},
	)
}

func (i initproject) generateIgnoreFile() error {
	log.Info("Generate Ignore File")
	return runn.Execute()
}

func (i initproject) generateTypicalWrapper() error {
	log.Info("Generate Typical Wrapper")
	return runn.Execute()
}
