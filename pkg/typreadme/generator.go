package typreadme

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typictx"
	"github.com/typical-go/typical-go/pkg/typienv"
)

// Generator is responsible to generate readme
type Generator struct{}

// GenerateReadme to generate readme
func (Generator) GenerateReadme(ctx *typictx.Context) (err error) {
	var file *os.File
	log.Infof("Generate Readme: %s", typienv.Readme)
	if file, err = os.Create(typienv.Readme); err != nil {
		return
	}
	defer file.Close()
	return readme(file, ctx)
}
