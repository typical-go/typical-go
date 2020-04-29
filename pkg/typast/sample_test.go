package typast_test

type sampleInterface interface {
	sampleMethod()
}

// sampleStruct
// @tag1
// @tag2 (key1 key2 key3=value3)
type sampleStruct struct {
	sampleInt    int
	sampleString string
}

func sampleFunction() {
	// intentionally blank
}
