package helloworld

// SomeInterface is just some interface
// [mock]
type SomeInterface interface {
	Method()
}

// SomeStruct is just some struct
type SomeStruct struct{}

// NewSomeStruct return new some struct
// [constructor]
func NewSomeStruct() *SomeStruct {
	return &SomeStruct{}
}
