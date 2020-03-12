package typcfg

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore"
)

const (
	defaultDotEnv = ".env"
)

// TypicalConfiguration of typical project
type TypicalConfiguration struct {
	loader      Loader
	configurers []Configurer
}

// New return new instance of Configuration
func New() *TypicalConfiguration {
	return &TypicalConfiguration{
		loader: &defaultLoader{},
	}
}

// WithLoader to set loader
func (c *TypicalConfiguration) WithLoader(loader Loader) *TypicalConfiguration {
	c.loader = loader
	return c
}

// AppendConfigurer to append configurer
func (c *TypicalConfiguration) AppendConfigurer(configurers ...Configurer) *TypicalConfiguration {
	c.configurers = append(c.configurers, configurers...)
	return c
}

// Loader of configuration
func (c *TypicalConfiguration) Loader() Loader {
	return c.loader
}

// Setup the configuration to be ready to use for the app and build-tool
func (c *TypicalConfiguration) Setup() (err error) {
	var (
		f *os.File
	)
	if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
		log.Infof("Generate new project environment at '%s'", defaultDotEnv)
		if f, err = os.Create(defaultDotEnv); err != nil {
			return
		}
		defer f.Close()
		if err = c.Write(f); err != nil {
			return
		}
	}
	// TODO: load env
	return
}

// Write typical configuration
func (c *TypicalConfiguration) Write(w io.Writer) (err error) {
	store := c.Store()
	for _, field := range store.Fields() {
		var v interface{}
		if field.IsZero {
			v = field.Default
		} else {
			v = field.Value
		}
		if _, err = fmt.Fprintf(w, "%s=%v\n", field.Name, v); err != nil {
			return
		}
	}
	return
}

// Store to return config store that contain config informatino
func (c *TypicalConfiguration) Store() *typcore.ConfigStore {
	store := typcore.NewConfigStore()
	for _, configurer := range c.configurers {
		cfg := configurer.Configure(c.loader)
		if cfg == nil {
			panic("Configure return nil detail")
		}
		fields := retrieveFields(cfg.Name, cfg.Spec)
		store.Put(cfg.Name, typcore.NewConfigBean(cfg.Name, fields, cfg.Constructor))
	}

	return store
}

func retrieveFields(prefix string, spec interface{}) (fields []*typcore.ConfigField) {
	val := reflect.Indirect(reflect.ValueOf(spec))
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !fieldIgnored(field) {
			name := fmt.Sprintf("%s_%s", prefix, fieldName(field))
			fields = append(fields, &typcore.ConfigField{
				Name:     name,
				Type:     field.Type.Name(),
				Default:  fieldDefault(field),
				Required: fieldRequired(field),
				Value:    val.Field(i).Interface(),
				IsZero:   val.Field(i).IsZero(),
			})
		}
	}
	return
}

func fieldRequired(field reflect.StructField) (required bool) {
	if v, ok := field.Tag.Lookup("required"); ok {
		required, _ = strconv.ParseBool(v)
	}
	return
}

func fieldIgnored(field reflect.StructField) (ignored bool) {
	if v, ok := field.Tag.Lookup("ignored"); ok {
		ignored, _ = strconv.ParseBool(v)
	}
	return
}

func fieldDefault(field reflect.StructField) string {
	return field.Tag.Get("default")
}

func fieldName(field reflect.StructField) (name string) {
	name = strings.ToUpper(field.Name)
	if v, ok := field.Tag.Lookup("envconfig"); ok {
		name = v
	}
	return
}
