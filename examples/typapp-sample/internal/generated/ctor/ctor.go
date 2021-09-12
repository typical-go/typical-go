package ctor

import (
	a "github.com/typical-go/typical-go/examples/typapp-sample/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

// DO NOT EDIT. This is code generated file

func init() {
	typapp.Provide("", a.HelloWorld)
	typapp.Provide("typical", a.HelloTypical)
}
