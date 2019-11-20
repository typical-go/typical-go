package app

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/app/common"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/utility/runn"
)

// InitProject iniate new project
func InitProject(parent, pkg string) error {
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
	return runn.Execute(
		common.Mkdir{Path: i.Path("cmd")},
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
