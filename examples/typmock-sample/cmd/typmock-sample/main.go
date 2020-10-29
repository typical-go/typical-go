package main

import (
	"os"

	"github.com/typical-go/typical-go/examples/typmock-sample/internal/app"
	"github.com/typical-go/typical-go/examples/typmock-sample/internal/greeter"
)

func main() {
	app.Start(os.Stdout, greeter.New())
}
