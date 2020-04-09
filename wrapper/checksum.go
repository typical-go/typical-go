package wrapper

import (
	"bytes"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Checksum contain hash of file
type Checksum struct {
	source string
	data   []byte
}

// CreateChecksum to create checksum from source
func CreateChecksum(source string) (c *Checksum, err error) {
	h := sha256.New()
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if b, err := ioutil.ReadFile(path); err == nil {
			h.Write(b)
		}
		return nil
	})
	return &Checksum{
		source: source,
		data:   h.Sum(nil),
	}, nil
}

// IsSame return true if filename contain same checksum
func (c *Checksum) IsSame(filename string) bool {
	if b, err := ioutil.ReadFile(filename); err == nil {
		return bytes.Compare(c.data, b) == 0
	}
	return false
}

// Save checksum to filename
func (c *Checksum) Save(filename string) error {
	return ioutil.WriteFile(filename, c.data, 0777)
}
