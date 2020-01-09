package tmpl

// TypicalwData is data for typicalw template
type TypicalwData struct {
	DescriptorFile    string
	ChecksumFile      string
	LayoutTemp        string
	BuildtoolMainPath string
	BuildtoolBin      string
}

// Typicalw template
const Typicalw = `#!/bin/bash
set -e

CHECKSUM_DATA=$(cksum {{.DescriptorFile}})

if ! [ -s {{.ChecksumFile}} ]; then
	mkdir -p {{.LayoutTemp}}
	cksum {{.DescriptorFile}} > {{.ChecksumFile}}
else
	CHECKSUM_UPDATED=$([ "$CHECKSUM_DATA" == "$(cat {{.ChecksumFile}} )" ] ; echo $?)
fi

if [ "$CHECKSUM_UPDATED" == "1" ] || ! [[ -f {{.BuildtoolBin}} ]] ; then 
	echo $CHECKSUM_DATA > .typical-tmp/checksum
	echo "Compile Typical-Build"
	go build -o {{.BuildtoolBin}} ./{{.BuildtoolMainPath}}
fi

./{{.BuildtoolBin}} $@`
