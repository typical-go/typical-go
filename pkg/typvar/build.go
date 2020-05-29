package typvar

import (
	"fmt"
	"os"
)

var (
	Bin           string
	Src           string
	BuildChecksum string
	BuildToolSrc  string
	BuildToolBin  string
)

func Init() {
	Bin = fmt.Sprintf("%s/bin", TypicalTmp)
	Src = fmt.Sprintf("%s/src", TypicalTmp)
	BuildChecksum = fmt.Sprintf("%s/checksum", TypicalTmp)
	BuildToolSrc = fmt.Sprintf("%s/build-tool", Src)
	BuildToolBin = fmt.Sprintf("%s/build-tool", Bin)

}

func Wrap(typicalTmp, projectPkg string) {
	TypicalTmp = typicalTmp
	ProjectPkg = projectPkg
	Init()

	os.MkdirAll(BuildToolSrc, 0777)
	os.MkdirAll(Bin, 0777)
}
