package typgo_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestTestCmd(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.TestCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Test the project",
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}
}

func TestStdTest(t *testing.T) {
	t.Run("predefined", func(t *testing.T) {
		stdtest := &typgo.StdTest{
			Timeout:      123 * time.Second,
			CoverProfile: "some-profile",
			Packages:     []string{"pkg1", "pkg2"},
		}
		require.Equal(t, "some-profile", stdtest.GetCoverProfile())
		require.Equal(t, 123*time.Second, stdtest.GetTimeout())
		require.Equal(t, []string{"pkg1", "pkg2"}, stdtest.GetPackages(nil))
	})
	t.Run("default", func(t *testing.T) {
		stdtest := &typgo.StdTest{}
		ctx := &typgo.Context{
			Descriptor: &typgo.Descriptor{
				Layouts: []string{"pkg3", "pkg4"},
			},
		}
		require.Equal(t, "cover.out", stdtest.GetCoverProfile())
		require.Equal(t, 25*time.Second, stdtest.GetTimeout())
		require.Equal(t, []string{"./pkg3/...", "./pkg4/..."}, stdtest.GetPackages(ctx))
	})

}
