package typbuildtool

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typobj"
	"github.com/urfave/cli/v2"
)

const (
	defaultDotEnv = ".env"
)

func (t buildtool) before(ctx *cli.Context) (err error) {
	if err = t.Validate(); err != nil {
		return
	}
	cfgFields := ConfigFields(t.Context)
	// if _, err = metadata.Update("config_fields", cfgFields); err != nil {
	// 	log.Fatal(err.Error())
	// }
	if err = GenerateEnvfile(cfgFields); err != nil {
		return
	}
	return
}

// ConfigFields return config list
func ConfigFields(ctx *typctx.Context) (fields []typobj.Field) {
	for _, module := range ctx.AllModule() {
		if configurer, ok := module.(typobj.Configurer); ok {
			prefix, spec, _ := configurer.Configure()
			fields = append(fields, typobj.Fields(prefix, spec)...)
		}
	}
	return
}

// GenerateEnvfile to generate .env file if not exist
func GenerateEnvfile(fields []typobj.Field) (err error) {
	if _, err = os.Stat(defaultDotEnv); !os.IsNotExist(err) {
		return
	}
	log.Infof("Generate new project environment at '%s'", defaultDotEnv)
	var file *os.File
	if file, err = os.Create(defaultDotEnv); err != nil {
		return
	}
	defer file.Close()
	for _, field := range fields {
		var v interface{}
		if field.IsZero {
			v = field.Default
		} else {
			v = field.Value
		}
		s := fmt.Sprintf("%s=%v\n", field.Name, v)
		file.WriteString(s)
	}
	return
}
