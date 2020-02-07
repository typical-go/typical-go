package tmpl

// TypicalwData is data for typicalw template
type TypicalwData struct {
	DescriptorPackage string
	DescriptorFile    string
	ChecksumFile      string
	LayoutTemp        string
}

// Typicalw template
const Typicalw = `#!/bin/bash

set -e

if ! [ -s .typical-tmp/build-tool/main.go ]; then
	mkdir -p .typical-tmp/build-tool
	echo "package main

import (
	\"github.com/typical-go/typical-go/pkg/typbuild\"
	\"{{.DescriptorPackage}}\"
)

func main() {
	typbuild.Run(&typical.Descriptor)
}" >> .typical-tmp/build-tool/main.go
fi

CHECKSUM_DATA=$(cksum {{.DescriptorFile}})

if ! [ -s {{.ChecksumFile}} ]; then
	mkdir -p {{.LayoutTemp}}
	echo $CHECKSUM_DATA > {{.ChecksumFile}}
else
	CHECKSUM_UPDATED=$([ "$CHECKSUM_DATA" == "$(cat {{.ChecksumFile}} )" ] ; echo $?)
fi

if [ "$CHECKSUM_UPDATED" == "1" ] || ! [[ -f ./.typical-tmp/bin/build-tool ]] ; then 
	echo $CHECKSUM_DATA > .typical-tmp/checksum
	echo "Compile Typical-Build"
	go build -o .typical-tmp/bin/build-tool .typical-tmp/build-tool/main.go
fi

./.typical-tmp/bin/build-tool $@`
