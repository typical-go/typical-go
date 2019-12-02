package typcfg

// Configurer responsible to create config
type Configurer interface {
	Configure() (prefix string, spec interface{}, loadFn interface{})
}

// IsConfigurer return true if object implementation of configurer
func IsConfigurer(obj interface{}) (ok bool) {
	_, ok = obj.(Configurer)
	return
}
