package app

const appSrc = `package app

import (
	"fmt"
	"{{.Pkg}}/app/config"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

// Module of application
func Module() interface{} {
	return &module{}
}

type module struct {}

func (module) Action() interface{} {
	return func(cfg config.Config) {
		fmt.Printf("Hello %s\n", cfg.Hello)
	}
}

func (module) Configure() (prefix string, spec, loadFn interface{}) {
	prefix = "APP"
	spec = &config.Config{}
	loadFn = func(loader typcfg.Loader) (cfg config.Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}
`

const appSrcTest = `package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typmodule"

	"{{.Pkg}}/app"
)

func TestModule(t *testing.T) {
	a := app.Module()
	require.True(t, typmodule.IsActionable(a))
	require.True(t, typcfg.IsConfigurer(a))
}
`

const configSrc = "package config\n\n// Config of app\ntype Config struct {\n	Hello string `default:\"World\"`\n}"

const ctxSrc = `package typical

import (
	"{{.Pkg}}/app"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Context of Project
var Context = &typctx.Context{
	Name:      "sample",
	Version:   "0.0.1",
	Package:   "sample",
	AppModule: app.Module(),
	Releaser: typrls.Releaser{
		Targets: []typrls.Target{"linux/amd64", "darwin/amd64"},
	},
	ConfigLoader: typcfg.DefaultLoader(),
}
`

const blankCtxSrc = `package typical

import (
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Context of Project
var Context = &typctx.Context{
	Name:      "sample",
	Version:   "0.0.1",
	Package:   "sample",
	Releaser: typrls.Releaser{
		Targets: []typrls.Target{"linux/amd64", "darwin/amd64"},
	},
}
`

const typicalw = `#!/bin/bash
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
	go build -tags typical -o $BIN/$PRE_BUILDER ./$CMD/$PRE_BUILDER
fi

./$BIN/$PRE_BUILDER $CHECKSUM_UPDATED
./$BIN/$BUILD_TOOL $@`

const gomod = `module {{.Pkg}}

go 1.12

require github.com/typical-go/typical-go v{{.TypicalVersion}}
`

const gitignore = `/bin
/release
/.typical-metadata
/vendor
.envrc
.env
*.test
*.out`

const moduleSrc = `package {{.Name}}

import "github.com/typical-go/typical-go/pkg/typcfg"

// Config is configuration of module mymodule
type Config struct {
	// TODO:
}

// Module of {{.Name}}
func Module() interface{} {
	return &module{}
}

type module struct {}


func (module) Configure() (prefix string, spec, loadFn interface{}) {
	prefix = "{{.Prefix}}"
	spec = &Config{}
	loadFn = func(loader typcfg.Loader) (cfg Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}

func (m *module) Provide() []interface{} {
	return []interface{}{
		// TODO: functions to be provided as dependency
	}
}

func (m *module) Prepare() []interface{} {
	return []interface{}{
		// TODO: functions to check/prepare the dependencies
	}
}

func (m *module) Destroy() []interface{} {
	return []interface{}{
		// TODO: functions to destroy dependencies
	}
}
`

const moduleSrcTest = `package {{.Name}}

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typmodule"
)

func TestModule(t *testing.T) {
	m := Module()
	require.True(t, typmodule.IsProvider(m))
	require.True(t, typmodule.IsDestroyer(m))
	require.True(t, typmodule.IsProvider(m))
	require.True(t, typcfg.IsConfigurer(m))
}
`
