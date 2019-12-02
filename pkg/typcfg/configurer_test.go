package typcfg_test

import "testing"

type SampleSpec struct {
	Field1 string `default:"hello" required:"true"`
	Field2 string `default:"world"`
	Field3 string `ignored:"true"`
	Field4 int    `envconfig:"ALIAS"`
}

type configurerImpl struct{}

func (configurerImpl) Configure() (prefix string, specFn interface{}) {
	return "TEST", func() SampleSpec {
		return SampleSpec{}
	}
}

func TestConfigurer(t *testing.T) {

}
