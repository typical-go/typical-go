package typapp_test

import (
	"fmt"
	"log"
	"strings"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"go.uber.org/dig"
)

func ExampleApplication() {
	typapp.Provide("", func() string { return "world" })

	application := &typapp.Application{
		StartFn: func(text string) {
			fmt.Printf("hello %s\n", text)
		},
		ShutdownFn: func() {
			fmt.Println("bye2")
		},
		ExitSigs: []syscall.Signal{syscall.SIGTERM, syscall.SIGINT},
	}

	if err := application.Run(); err != nil {
		log.Fatal(err.Error())
	}

	// Output: hello world
	// bye2
}

func TestRun(t *testing.T) {
	defer typapp.Reset()

	var out strings.Builder
	typapp.Provide("", func() string { return "success" })

	application := typapp.Application{
		StartFn: func(s string) {
			fmt.Fprintln(&out, s)
		},
		ShutdownFn: func() {
			fmt.Fprintln(&out, "clean")
		},
	}
	require.NoError(t, application.Run())
	require.Equal(t, "success\nclean\n", out.String())
}

func TestRun_BadStartFn(t *testing.T) {
	defer typapp.Reset()

	application := typapp.Application{
		StartFn:    "bad-start-fn",
		ShutdownFn: "bad-shutdown-fn",
	}
	require.EqualError(t, application.Run(), "can't invoke non-function bad-start-fn (type string); can't invoke non-function bad-shutdown-fn (type string)")
}

func TestRun_BadConstructor(t *testing.T) {
	defer typapp.Reset()
	typapp.Provide("", "bad-constructor")

	application := typapp.Application{
		StartFn: func() {},
	}
	require.EqualError(t, application.Run(), "must provide constructor function, got bad-constructor (type string)")
}

func TestRun_ProvideDigContainer(t *testing.T) {
	defer typapp.Reset()

	var out strings.Builder
	typapp.Provide("", func() string { return "success" })

	application := typapp.Application{
		StartFn: func(di *dig.Container) error {
			return di.Invoke(func(s string) {
				fmt.Fprintln(&out, s)
			})
		},
		ShutdownFn: func() {
			fmt.Fprintln(&out, "clean")
		},
	}
	require.NoError(t, application.Run())
	require.Equal(t, "success\nclean\n", out.String())
}
