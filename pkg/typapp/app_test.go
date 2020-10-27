package typapp_test

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func ExampleStart() {
	typapp.Provide("", func() string { return "World" })

	start := func(text string) { fmt.Printf("Hello %s\n", text) }
	shutdown := func() { fmt.Println("clean something") }

	if err := typapp.Run(start, shutdown); err != nil {
		log.Fatal(err.Error())
	}

	// Output: Hello World
	// clean something
}

func TestStart(t *testing.T) {
	defer typapp.Reset()

	var out strings.Builder
	typapp.Provide("", func() string { return "success" })
	startFn := func(s string) { fmt.Fprintln(&out, s) }
	shutdownFn := func() { fmt.Fprintln(&out, "clean") }

	require.NoError(t, typapp.Run(startFn, shutdownFn))
	require.Equal(t, "success\nclean\n", out.String())
}

func fnErr(errMsg string) func() error {
	return func() error {
		return errors.New(errMsg)
	}
}
