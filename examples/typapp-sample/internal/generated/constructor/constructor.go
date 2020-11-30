package constructor

/* DO NOT EDIT. This file generated due to '@ctor' annotation*/

import (
	a "github.com/typical-go/typical-go/examples/typapp-sample/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() {
	typapp.Provide("", a.HelloWorld)
	typapp.Provide("typical", a.HelloTypical)
}
