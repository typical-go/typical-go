package constructor

/* Autogenerated by Typical-Go. DO NOT EDIT.

TagName:
	@ctor

Help:
	https://pkg.go.dev/github.com/typical-go/typical-go/pkg/typapp?tab=doc#CtorAnnotation
*/

import (
	a "github.com/typical-go/typical-go/examples/typmock-sample/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() {
	typapp.Provide("", a.GetWriter)
	typapp.Provide("", a.NewGreeter)
}