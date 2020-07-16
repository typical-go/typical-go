package typapp_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestApp_Run(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		app := &typapp.App{
			EntryPoint: fnErr("entry-point-err"),
			Dtors: []*typapp.Destructor{
				{Fn: fnErr("dtor-1-err")},
				{Fn: fnErr("dtor-2-err")},
			},
		}
		require.EqualError(t, app.Run(), "entry-point-err; dtor-1-err; dtor-2-err")
	})
	t.Run("bad constructor", func(t *testing.T) {
		app := &typapp.App{
			Ctors: []*typapp.Constructor{
				{Fn: "bad-contructor"},
			},
		}
		require.EqualError(t, app.Run(), "must provide constructor function, got bad-contructor (type string)")
	})
	t.Run("provide constructor in entry-point", func(t *testing.T) {
		var text string
		app := &typapp.App{
			EntryPoint: func(s string) { text = s },
			Ctors: []*typapp.Constructor{
				{Fn: func() string { return "success" }},
			},
		}
		require.NoError(t, app.Run())
		require.Equal(t, "success", text)
	})

}

func fnErr(errMsg string) func() error {
	return func() error {
		return errors.New(errMsg)
	}
}
