package typbuildtool

import (
	"fmt"
	"io"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// WriteEnv to write env file
func WriteEnv(w io.Writer, configMap typcore.ConfigMap) (err error) {
	for _, cfgDetail := range configMap {
		var v interface{}
		if cfgDetail.IsZero {
			v = cfgDetail.Default
		} else {
			v = cfgDetail.Value
		}
		if _, err = fmt.Fprintf(w, "%s=%v\n", cfgDetail.Name, v); err != nil {
			return
		}
	}
	return
}
