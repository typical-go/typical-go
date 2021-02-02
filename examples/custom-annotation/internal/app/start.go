package app

import (
	"fmt"
)

// Start application
func Start() {
	fmt.Println("Hello World")
}

type (
	// MyStruct1 ...
	// @mytag
	MyStruct1 struct{}
	// MyStruct2 ...
	// @mytag (field1:"value1")
	MyStruct2 struct{}
	// MyIntf1 ...
	// @mytag
	MyIntf1 interface{}
	// MyIntf2 ...
	// @mytag
	MyIntf2 interface{}
)

// MyFunc1 ...
// @mytag
func MyFunc1() {}

// MyFunc2 ...
// @mytag
func MyFunc2() {}
