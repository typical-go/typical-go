package common

import "fmt"

// KeyString short from parameter
type KeyString struct {
	Key    string
	String string
}

// SimpleFormat return string of key-string with simple format
func (k KeyString) SimpleFormat(sep string) string {
	return fmt.Sprintf("%s%s%s", k.Key, sep, k.String)
}

// Format return string of key-string
func (k KeyString) Format(fn func(key, s string) string) string {
	return fn(k.Key, k.String)
}
