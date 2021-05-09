package typapp_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"go.uber.org/dig"
)

func ExampleStartApp() {
	typapp.Reset() // make sure constructor and container is empty (optional)
	typapp.Provide("", func() string { return "world" })

	startFn := func(text string) { fmt.Printf("hello %s\n", text) }
	stopFn := func() { fmt.Println("bye2") }

	if err := typapp.StartApp(startFn, stopFn); err != nil {
		log.Fatal(err)
	}

	// Output: hello world
	// bye2
}

func ExampleInvoke() {
	typapp.Reset()

	typapp.Provide("", func() string { return "world" })

	err := typapp.Invoke(func(text string) {
		fmt.Printf("hello %s\n", text)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output: hello world
}

func ExampleInvoke_second() {
	typapp.Reset()

	typapp.Provide("t1", func() string { return "hello" }) // provide same type
	typapp.Provide("t2", func() string { return "world" }) // provide same type

	type param struct {
		dig.In
		Text1 string `name:"t1"`
		Text2 string `name:"t2"`
	}

	printHello := func(p param) {
		fmt.Printf("%s %s\n", p.Text1, p.Text2)
	}

	if err := typapp.Invoke(printHello); err != nil {
		log.Fatal(err)
	}

	// Output: hello world
}

func TestRun(t *testing.T) {
	defer typapp.Reset()

	var out strings.Builder
	typapp.Provide("", func() string { return "success" })

	startFn := func(s string) { fmt.Fprintln(&out, s) }
	stopFn := func() { fmt.Fprintln(&out, "clean") }

	require.NoError(t, typapp.StartApp(startFn, stopFn))
	require.Equal(t, "success\nclean\n", out.String())
}

func TestSet(t *testing.T) {
	expectedConstructors := []*typapp.Constructor{}
	expectedContainer := dig.New()
	typapp.SetConstructors(expectedConstructors)
	typapp.SetContainer(expectedContainer)

	require.Equal(t, expectedConstructors, typapp.Constructors())

	container, err := typapp.Container()
	require.NoError(t, err)
	require.Equal(t, expectedContainer, container)
}

func TestRun_BadStartFn(t *testing.T) {
	defer typapp.Reset()
	startFn := "bad-start-fn"
	stopFn := "bad-shutdown-fn"
	require.EqualError(t, typapp.StartApp(startFn, stopFn),
		"can't invoke non-function bad-start-fn (type string); can't invoke non-function bad-shutdown-fn (type string)")
}

func TestRun_BadConstructor(t *testing.T) {
	defer typapp.Reset()
	typapp.Provide("", "bad-constructor")

	require.EqualError(t, typapp.StartApp(func() {}, nil), "must provide constructor function, got bad-constructor (type string)")
}

func TestRun_ProvideDigContainer(t *testing.T) {
	defer typapp.Reset()

	var out strings.Builder
	typapp.Provide("", func() string { return "success" })

	invokeFn := func(di *dig.Container) error {
		return di.Invoke(func(s string) {
			fmt.Fprintln(&out, s)
		})
	}

	require.NoError(t, typapp.Invoke(invokeFn))
	require.Equal(t, "success\n", out.String())
}
