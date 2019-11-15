package app

import "github.com/typical-go/typical-go/pkg/typictx"

// Module of application
func Module() typictx.AppModule {
	return applicationModule{}
}

type applicationModule struct {
}

func (m applicationModule) Run() interface{} {
	return start
}
