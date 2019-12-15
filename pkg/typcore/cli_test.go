package typcore_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewCli(t *testing.T) {
	ctx := &typcore.Context{}
	obj := struct{}{}
	cli := typcore.NewCli(ctx, obj)
	require.Equal(t, ctx, cli.Context())
	require.Equal(t, obj, cli.Object())
}
