package typcore

import "os"

// TempFolder contain temporary information for typical build
type TempFolder string

// BuildToolBin is build-tool binary path
func (t TempFolder) BuildToolBin() string {
	return string(t) + "/bin/build-tool"
}

// BuildToolSrc is build-tool source path
func (t TempFolder) BuildToolSrc() string {
	return string(t) + "/build-tool/main.go"
}

// Checksum is checksum path
func (t TempFolder) Checksum() string {
	return string(t) + "/checksum"
}

// Mkdir to make temp folderif not exist
func (t TempFolder) Mkdir() {
	os.MkdirAll(string(t)+"/build-tool", os.ModePerm)
	os.MkdirAll(string(t)+"/bin", os.ModePerm)
}
