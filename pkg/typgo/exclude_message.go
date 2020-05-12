package typgo

import (
	"strings"

	"github.com/typical-go/typical-go/pkg/typvar"
)

// ExcludeMessage return true is message mean to be exclude
func ExcludeMessage(msg string) bool {
	msg = strings.ToLower(msg)
	for _, prefix := range typvar.Rls.ExclMsgPrefix {
		if strings.HasPrefix(msg, strings.ToLower(prefix)) {
			return true
		}
	}
	return false
}
