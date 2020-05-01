package typfactory_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/typfactory"
)

func TestTypicalw(t *testing.T) {
	testWriter(t,
		testcase{
			Writer: &typfactory.Typicalw{
				TypicalSource: "some-source",
				TypicalTmp:    "some-tmp",
			},
			expected: `#!/bin/bash

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
		testcase{
			Writer: &typfactory.Typicalw{
				TypicalSource: "some-source",
				TypicalTmp:    "some-tmp",
				ProjectPkg:    "some-project-pkg",
			},
			expected: `#!/bin/bash

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
	)

}
