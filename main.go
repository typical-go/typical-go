package main

import (
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	typapp.Start(app.Start)
}
