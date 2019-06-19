package stmt

import (
	"github.com/typical-go/typical-go/appcontext"

	"encoding/json"
	"io/ioutil"
)

// SaveContext save context to corresponding source
func SaveContext(context appcontext.Context, source string) Statement {
	return &saveContext{
		source:  source,
		context: context,
	}
}

type saveContext struct {
	source  string
	context appcontext.Context
}

func (c saveContext) Run() error {
	b, _ := json.MarshalIndent(c.context, "", "    ")
	return ioutil.WriteFile(c.source, b, 0644)
}
