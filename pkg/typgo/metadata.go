package typgo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	// Metadata is simple file-based json database
	Metadata struct {
		Path  string                 `json:"-"`
		Extra map[string]interface{} `json:"extra"`
	}
)

// OpenMetadata to open metadata
func OpenMetadata(path string) (*Metadata, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return createEmpty(path)
	}

	var (
		metadata Metadata
		err      error
	)

	b, _ := ioutil.ReadFile(path)
	if err = json.Unmarshal(b, &metadata); err != nil {
		return nil, err
	}
	metadata.Path = path

	return &metadata, nil
}

func createEmpty(path string) (db *Metadata, err error) {
	db = &Metadata{
		Path:  path,
		Extra: map[string]interface{}{},
	}
	err = db.Save()
	return
}

// Save db to file
func (d *Metadata) Save() (err error) {
	b, _ := json.Marshal(d)
	return ioutil.WriteFile(d.Path, b, 0777)
}
