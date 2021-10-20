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
	Coders  []Coder
	Comment string
)

func WriteCoder(c *typgo.Context, coder Coder, target string) error {
	defer typgo.GoImports(c, target)

	dir := filepath.Dir(target)
	os.MkdirAll(dir, 0777)

	data := coder.Code()
	return ioutil.WriteFile(target, []byte(data), 0777)
}

//
// Coders
//

var _ Coder = (Coders)(nil)

func (s Coders) Code() string {
	var b strings.Builder
	for _, src := range s {
		b.WriteString(src.Code())
		b.WriteString("\n")
	}

	return b.String()
}

//
// Comment
//

var _ Coder = (Coders)(nil)

func (c Comment) Code() string {
	return "// " + string(c)
}
