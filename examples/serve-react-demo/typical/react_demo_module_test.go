package typical_test

import (
	"testing"

	"github.com/typical-go/typical-go/examples/serve-react-demo/typical"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func TestNewReactDemoModule(t *testing.T) {
	t.Run("SHOULD implement builder", func(t *testing.T) {
		var _ typbuildtool.Builder = &typical.ReactDemoModule{}
	})
	t.Run("SHOULD implement cleaner", func(t *testing.T) {
		var _ typbuildtool.Cleaner = &typical.ReactDemoModule{}
	})
}
