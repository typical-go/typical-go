package generated

import (
	a "github.com/typical-go/typical-go/examples/typapp-sample/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

// DO NOT EDIT. Code-generated file.
func init() {
	// TODO: provide struct constructor
	// TODO: provide interface constructor
	typapp.Provide("", a.HelloWorld)
	typapp.Provide("typical", a.HelloTypical)
}
