package typcfg

// Configurer responsible to create config
type Configurer interface {
	Configure() []*Configuration
}

// Field of config
type Field struct {
	Name     string
	Type     string
	Default  string
	Value    interface{}
	IsZero   bool
	Required bool
}
