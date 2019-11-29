// +build typical

package main

import (
	_ "github.com/typical-go/typical-go/internal/dependency"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/typical"
)

func main() {
	typbuildtool.Run(typical.Context)
}
