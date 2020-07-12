package typtmpl_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestTypicalw(t *testing.T) {
	typtmpl.TestTemplate(t, []typtmpl.TestCase{
		{
			Template: &typtmpl.Typicalw{
				TypicalSource: "some-source",
				TypicalTmp:    "some-tmp",
			},
			Expected: `#!/bin/bash

set -e

TYPSRC=some-source
TYPTMP=some-tmp
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	go build -o $TYPGO $TYPSRC
fi

$TYPGO wrap \
	-typical-tmp=$TYPTMP \

$TYPTMP/bin/build-tool $@
`,
		},
		{
			Template: &typtmpl.Typicalw{
				TypicalSource: "some-source",
				TypicalTmp:    "some-tmp",
				ProjectPkg:    "some-project-pkg",
			},
			Expected: `#!/bin/bash

set -e

TYPSRC=some-source
TYPTMP=some-tmp
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	go build -o $TYPGO $TYPSRC
fi

$TYPGO wrap \
	-typical-tmp=$TYPTMP \
	-project-pkg="some-project-pkg" \

$TYPTMP/bin/build-tool $@
`,
		},
	})

}
