package typcfg

// Configurer responsible to create config
type Configurer interface {
	Configurations() []*Configuration
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

// GetValue to get value or default value if no value
func (f *Field) GetValue() interface{} {
	if f.IsZero {
		return f.Default
	}
	return f.Value
}
