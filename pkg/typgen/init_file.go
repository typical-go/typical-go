package typgen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	InitFile struct {
		AliasGen *AliasGenerator
		InitBody []string
	}
)

var (
	defaultInitFileImports = map[string]string{
		"github.com/typical-go/typical-go/pkg/typapp": "",
	}
)

func NewInitFile() *InitFile {
	return &InitFile{
		AliasGen: NewAliasGenerator(defaultInitFileImports),
	}
}

func (i *InitFile) WriteTo(c *typgo.Context, target string) error {
	defer typgo.GoImports(c, target)

	dir := filepath.Dir(target)
	os.MkdirAll(dir, 0777)

	sourceCoders := Coders{
		&File{
			Name:   filepath.Base(dir),
			Import: i.AliasGen.Imports(),
		},
		Comment("DO NOT EDIT. Code-generated file."),
		&Function{
			Name: "init",
			Body: i.InitBody,
		},
	}

	data := sourceCoders.Code()
	return ioutil.WriteFile(target, []byte(data), 0777)
}

func (i *InitFile) PutInit(s string) {
	i.InitBody = append(i.InitBody, s)
}

func (i *InitFile) PutInitSprintf(format string, args ...interface{}) {
	i.PutInit(fmt.Sprintf(format, args...))
}

func (i *InitFile) ProvideConstructor(name, importPath, constructor string) {
	alias := i.AliasGen.Generate(importPath)
	i.PutInitSprintf(`typapp.Provide("%s", %s.%s)`, name, alias, constructor)
}
