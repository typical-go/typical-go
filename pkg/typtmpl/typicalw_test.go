package typtmpl_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestTypicalw(t *testing.T) {
	typtmpl.TestTemplate(t, []typtmpl.TestCase{
		{
			Template: &typtmpl.Typicalw{
				Src:        "some-src",
				TypicalTmp: "some-tmp",
				ProjectPkg: "some-project-pkg",
			},
			Expected: `#!/bin/bash

set -e

TYPTMP=some-tmp
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	echo "Build typical-go"
	go build -o $TYPGO github.com/typical-go/typical-go
fi

$TYPGO \
	-src="some-src" \
	-project-pkg="some-project-pkg" \
	-typical-tmp=$TYPTMP \
	$@
`,
		},
	})

}
