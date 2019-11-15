package prebuilder

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typictx"
	"github.com/typical-go/typical-go/pkg/typimodule"
)

const (
	defaultDotEnv = ".env"
)

// ConfigFields return config list
func ConfigFields(ctx *typictx.Context) (fields []typimodule.ConfigField) {
	for _, module := range ctx.AllModule() {
		if configurer, ok := module.(typimodule.Configurer); ok {
			fields = append(fields, configurer.Configure().ConfigFields()...)
		}
	}
	return
}

// GenerateEnvfile to generate .env file if not exist
func GenerateEnvfile(ctx *typictx.Context) (err error) {
	if _, err = os.Stat(defaultDotEnv); !os.IsNotExist(err) {
		return
	}
	log.Infof("Generate new project environment at '%s'", defaultDotEnv)
	var file *os.File
	if file, err = os.Create(defaultDotEnv); err != nil {
		return
	}
	defer file.Close()
	for _, field := range ConfigFields(ctx) {
		s := fmt.Sprintf("%s=%s\n", field.Name, field.Default)
		file.WriteString(s)
	}
	return
}
