package typgo_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestPrepareContext(t *testing.T) {
	var out strings.Builder
	c := typgo.NewPrepareContext(&typgo.Descriptor{ProjectName: "some-project"}, "some-state")
	c.Stdout = &out
	c.Info("some-information")
	c.Infof("I love you %d\n", 3000)
	require.Equal(t, "some-project:some-state> some-information\nsome-project:some-state> I love you 3000\n", out.String())
}
