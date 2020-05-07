package typast_test

type sampleInterface interface {
	sampleMethod()
}

// sampleStruct
// @tag1
// @tag2 {"key1":"", "key2": "", "key3":"value3"}
type sampleStruct struct {
	sampleInt    int
	sampleString string
}

func sampleFunction() {
	// intentionally blank
}

// GetWriter to get writer to greet the world
// @constructor
func sampleFunction2() {
	// intentionally blank
}

type (
	// @tag3
	sampleInterface2 interface {
		sampleMethod2()
	}

	// sampleStruct2 asdf
	// @tag4
	sampleStruct2 struct {
	}
)
