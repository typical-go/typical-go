package typtmpl_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

func TestTypicalw(t *testing.T) {

	testcases := []struct {
		TestName string
		typtmpl.Template
		Expected string
	}{
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

$TYPGO run \
	-src="some-src" \
	-project-pkg="some-project-pkg" \
	-typical-tmp=$TYPTMP \
	$@
`,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			var out strings.Builder
			require.NoError(t, tt.Execute(&out))
			require.Equal(t, tt.Expected, out.String())
		})
	}

}
