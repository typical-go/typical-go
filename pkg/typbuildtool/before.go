package typbuildtool

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

const (
	defaultDotEnv = ".env"
)

func (t buildtool) before(ctx *cli.Context) (err error) {
	var (
		f *os.File
	)
	if err = t.Validate(); err != nil {
		return
	}
	if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
		log.Infof("Generate new project environment at '%s'", defaultDotEnv)
		if f, err = os.Create(defaultDotEnv); err != nil {
			return
		}
		defer f.Close()
		_, configMap := typcore.CreateConfigMap(t.ProjectDescriptor)
		for _, field := range configMap {
			var v interface{}
			if field.IsZero {
				v = field.Default
			} else {
				v = field.Value
			}
			f.WriteString(fmt.Sprintf("%s=%v\n", field.Name, v))
		}
	}
	return
}
