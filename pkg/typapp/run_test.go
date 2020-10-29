package typapp_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func ExampleRun() {
	typapp.Provide("", func() string { return "world" })

	startFn := func(text string) { fmt.Printf("hello %s\n", text) }
	shutdownFn := func() { fmt.Println("bye2") }

	if err := typapp.Run(startFn, shutdownFn); err != nil {
		log.Fatal(err.Error())
	}

	// Output: hello world
	// bye2
}

func TestRun(t *testing.T) {
	defer typapp.Reset()

	var out strings.Builder
	typapp.Provide("", func() string { return "success" })

	startFn := func(s string) { fmt.Fprintln(&out, s) }
	shutdownFn := func() { fmt.Fprintln(&out, "clean") }

	require.NoError(t, typapp.Run(startFn, shutdownFn))
	require.Equal(t, "success\nclean\n", out.String())
}

func TestRun_BadStartFn(t *testing.T) {
	defer typapp.Reset()

	err := typapp.Run("bad-start-fn", "bad-shutdown-fn")
	require.EqualError(t, err, "can't invoke non-function bad-start-fn (type string); can't invoke non-function bad-shutdown-fn (type string)")
}

func TestRun_BadConstructor(t *testing.T) {
	defer typapp.Reset()
	typapp.Provide("", "bad-constructor")

	err := typapp.Run(func() {}, nil)
	require.EqualError(t, err, "must provide constructor function, got bad-constructor (type string)")
}
