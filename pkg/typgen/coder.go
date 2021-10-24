package typgen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	Coder interface {
		Code() string
	}
	Coders    []Coder
	CodeLine  string
	CodeLines []string
)

var (
	_ Coder = (Coders)(nil)
	_ Coder = (CodeLine)("")
)

func WriteCoder(c *typgo.Context, coder Coder, target string) error {
	defer typgo.GoImports(c, target)

	dir := filepath.Dir(target)
	os.MkdirAll(dir, 0777)

	data := coder.Code()
	return ioutil.WriteFile(target, []byte(data), 0777)
}

func (s Coders) Code() string {
	var b strings.Builder
	for _, src := range s {
		b.WriteString(src.Code())
		b.WriteString("\n")
	}

	return b.String()
}

func (c CodeLine) Code() string {
	return string(c)
}

func (c CodeLines) Code() string {
	var b strings.Builder
	for _, s := range c {
		b.WriteString(s)
		b.WriteString("\n")
	}
	return b.String()
}
