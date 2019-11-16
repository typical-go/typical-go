package typctx_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typctx"
)

func TestNewAppModule(t *testing.T) {
	fn := struct{}{}
	appModule := typctx.NewAppModule(fn)
	require.Equal(t, fn, appModule.Run())
}
