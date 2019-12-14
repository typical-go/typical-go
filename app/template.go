package app

const typicalw = `#!/bin/bash
set -e

CHECKSUM_DATA=$(cksum {{.ContextFile}})

if ! [ -s {{.ChecksumFile}} ]; then
	mkdir -p {{.LayoutMetadata}}
	cksum typical/context.go > {{.ChecksumFile}}
else
	CHECKSUM_UPDATED=$([ "$CHECKSUM_DATA" == "$(cat {{.ChecksumFile}} )" ] ; echo $?)
fi

if [ "$CHECKSUM_UPDATED" == "1" ] || ! [[ -f bin/typical-go-buildtool ]] ; then 
	echo $CHECKSUM_DATA > .typical-metadata/checksum
	echo "Compile Typical-Build"
	go build -o bin/typical-go-buildtool ./cmd/typical-go-buildtool
fi

./{{.BuildtoolBin}} $@`

const gomod = `module {{.Pkg}}

go 1.13

require github.com/typical-go/typical-go v{{.TypicalVersion}}
`

const gitignore = `/bin
/release
/.typical-metadata
/vendor
.envrc
.env
*.test
*.out`
