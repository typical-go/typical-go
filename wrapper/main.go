package wrapper

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typlog"
)

// Main function to run the typical-go
func Main(d *typgo.Descriptor) (err error) {
	wrapper := wrapper{
		Descriptor: d,
		Logger:     typlog.Logger{Name: "WRAPPER"},
	}

	return wrapper.app().Run(os.Args)
}
