package main

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/internal/app"
)

func main() {
	if err := app.Main(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
