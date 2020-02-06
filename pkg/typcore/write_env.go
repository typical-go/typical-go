package typcore

import (
	"fmt"
	"io"
)

// WriteEnv to write env file
func WriteEnv(w io.Writer, keys []string, configMap ConfigMap) (err error) {
	for _, key := range keys {
		var (
			v         interface{}
			cfgDetail = configMap[key]
		)
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
