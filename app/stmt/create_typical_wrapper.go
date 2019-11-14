package stmt

import (
	"io/ioutil"
)

var typicalwScript = []byte(`#!/bin/bash

BIN=${TYPICAL_BIN:-bin}
CMD=${TYPICAL_CMD:-cmd}
NAME=${TYPICAL_NAME:-typical-dev-tool}

go run ./$CMD/$NAME/* $@
`)

// CreateTypicalWrapper to create Typical Context in Target file
type CreateTypicalWrapper struct {
	Target string
}

// Run the create typical context
func (c CreateTypicalWrapper) Run() error {
	return ioutil.WriteFile(c.Target, typicalwScript, 0755)

}
