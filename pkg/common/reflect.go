package common

import (
	"reflect"
	"strings"
)

// IsNil to check if interface is nil
func IsNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}

// PackageName return package name of the interface
func PackageName(v interface{}) string {
	if IsNil(v) {
		return ""
	}
	s := reflect.TypeOf(v).String()
	if dot := strings.Index(s, "."); dot > 0 {
		if strings.HasPrefix(s, "*") {
			return s[1:dot]
		}
		return s[:dot]
	}
	return ""
}
