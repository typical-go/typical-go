package typapp_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typlog"
)

func TestTypicalApp_Preconditon(t *testing.T) {
	t.Run("GIVEN false precondition", func(t *testing.T) {
		var debugger strings.Builder
		defer typlog.SetOutput(&debugger)()

		c := &typbuildtool.BuildContext{
			Context: &typbuildtool.Context{
				Context: &typcore.Context{
					Descriptor: &typcore.Descriptor{},
				},
			},
		}

		require.NoError(t, typapp.
			AppModule(nil).
			WithPrecondition(false).
			Precondition(c))

		require.Equal(t, "[TYPICAL][INFO] Skip Precondition for typical-app\n", debugger.String())
	})
}

func TestConfigConstructor(t *testing.T) {
	expected := `func() (cfg *typapp_test.config, err error){
		cfg = new(typapp_test.config)
		if err = typcfg.Process("NAME", cfg); err != nil {
			return nil, err
		}
		return  
	}`

	require.Equal(t, expected, typapp.ConfigContructor(&typcore.Configuration{
		Name: "NAME",
		Spec: &config{},
	}))

}

type config struct{}
