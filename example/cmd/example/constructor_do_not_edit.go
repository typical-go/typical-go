package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/example/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() {
	typapp.AppendConstructor(
		app.NewSomeStruct,
	)
}
