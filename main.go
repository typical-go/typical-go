package main

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	if err := typapp.Run(app.Main); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
