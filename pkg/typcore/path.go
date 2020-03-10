package typcore

import (
	"fmt"
	"os"
)

// BuildToolBin is build-tool binary path
func BuildToolBin(tmp string) string {
	return tmp + "/bin/build-tool"
}

// BuildToolSrc is build-tool source path
func BuildToolSrc(tmp string) string {
	return tmp + "/build-tool/main.go"
}

// Checksum is checksum path
func Checksum(tmp string) string {
	return tmp + "/checksum"
}

// TypicalPackage is package of typical
func TypicalPackage(projectPackage string) string {
	return fmt.Sprintf("%s/typical", projectPackage)
}

// MakeTempDir to make temp folderif not exist
func MakeTempDir(tmp string) {
	os.MkdirAll(tmp+"/build-tool", os.ModePerm)
	os.MkdirAll(tmp+"/bin", os.ModePerm)
}
