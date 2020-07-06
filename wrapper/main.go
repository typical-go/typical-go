package wrapper

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typlog"
)

// Main function to run the typical-go
func Main() (err error) {
	wrapper := wrapper{
		Name:    typapp.Name,
		Version: typapp.Version,
		Logger:  typlog.Logger{Name: "WRAPPER"},
	}

	return wrapper.app().Run(os.Args)
}
