package tmpl

// Typicalw template
const Typicalw = `#!/bin/bash
set -e

CHECKSUM_DATA=$(cksum {{.DescriptorFile}})

if ! [ -s {{.ChecksumFile}} ]; then
	mkdir -p {{.LayoutMetadata}}
	cksum {{.DescriptorFile}} > {{.ChecksumFile}}
else
	CHECKSUM_UPDATED=$([ "$CHECKSUM_DATA" == "$(cat {{.ChecksumFile}} )" ] ; echo $?)
fi

if [ "$CHECKSUM_UPDATED" == "1" ] || ! [[ -f {{.BuildtoolBin}} ]] ; then 
	echo $CHECKSUM_DATA > .typical-metadata/checksum
	echo "Compile Typical-Build"
	go build -o {{.BuildtoolBin}} ./{{.BuildtoolMainPath}}
fi

./{{.BuildtoolBin}} $@`
