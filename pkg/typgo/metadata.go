package typgo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	// Metadata is simple file-based json database
	Metadata struct {
		Path   string
		Extras map[string]interface{}
	}
)

// OpenMetadata to open metadata
func OpenMetadata(path string) (db *Metadata, err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return createEmpty(path)
	}

	var (
		b []byte
		m map[string]interface{}
	)

	b, _ = ioutil.ReadFile(path)
	if err = json.Unmarshal(b, &m); err != nil {
		return
	}

	return &Metadata{
		Path:   path,
		Extras: m,
	}, nil
}

func createEmpty(path string) (db *Metadata, err error) {
	db = &Metadata{
		Path:   path,
		Extras: map[string]interface{}{},
	}
	err = db.Save()
	return
}

// Save db to file
func (d *Metadata) Save() (err error) {
	b, _ := json.Marshal(d.Extras)
	return ioutil.WriteFile(d.Path, b, 0777)
}
