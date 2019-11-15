package typictx_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typictx"
)

func TestNewAppModule(t *testing.T) {
	fn := struct{}{}
	appModule := typictx.NewAppModule(fn)
	require.Equal(t, fn, appModule.Run())
}
