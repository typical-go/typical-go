package typapp

var (
	appCtors []interface{}
)

// AppendConstructor to append constructor
func AppendConstructor(fn ...interface{}) {
	appCtors = append(appCtors, fn...)
}
