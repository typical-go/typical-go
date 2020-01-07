package typcore

// EntryPointer responsible to handle entry point
type EntryPointer interface {
	EntryPoint() interface{}
}

// Provider responsible to provide dependency
type Provider interface {
	Provide() []interface{}
}

// Preparer responsible to prepare
type Preparer interface {
	Prepare() []interface{}
}

// Destroyer responsible to destruct dependency
type Destroyer interface {
	Destroy() []interface{}
}

// Configurer responsible to create config
// `Prefix` is used by ConfigLoader to retrieve configuration value
// `Spec` (Specification) is used readme/env file generator. The value of spec will act as local environment value defined in .env file.
// `LoadFn` (Load Function) is required to provide in dependecies-injection container
type Configurer interface {
	Configure() (prefix string, spec interface{}, loadFn interface{})
}

// IsProvider return true if object implementation of provider
func IsProvider(obj interface{}) (ok bool) {
	_, ok = obj.(Provider)
	return
}

// IsPreparer return true obj implement Preparer
func IsPreparer(obj interface{}) (ok bool) {
	_, ok = obj.(Preparer)
	return
}

// IsDestroyer return true if object implementation of destructor
func IsDestroyer(obj interface{}) (ok bool) {
	_, ok = obj.(Destroyer)
	return
}

// IsConfigurer return true if object implementation of configurer
func IsConfigurer(obj interface{}) (ok bool) {
	_, ok = obj.(Configurer)
	return
}
