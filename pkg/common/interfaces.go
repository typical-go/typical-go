package common

// Interfaces is slice of interface{}
type Interfaces []interface{}

// Append item
func (i *Interfaces) Append(item ...interface{}) *Interfaces {
	*i = append(*i, item...)
	return i
}

// Slice of interface
func (i *Interfaces) Slice() []interface{} {
	return *i
}
