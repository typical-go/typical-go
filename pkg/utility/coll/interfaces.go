package coll

// Interfaces is slice of interface{}
type Interfaces []interface{}

// Append item
func (i Interfaces) Append(item ...interface{}) Interfaces {
	ptr := &i
	*ptr = append(*ptr, item...)
	return *ptr
}
