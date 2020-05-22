package typgo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	// Metadata is simple file-based json database
	Metadata struct {
		path string
		m    map[string]interface{}
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
		path: path,
		m:    m,
	}, nil
}

func createEmpty(path string) (db *Metadata, err error) {
	db = &Metadata{
		path: path,
		m:    map[string]interface{}{},
	}
	err = db.Save()
	return
}

// Save db to file
func (d *Metadata) Save() (err error) {
	b, _ := json.Marshal(d.m)
	return ioutil.WriteFile(d.path, b, 0777)
}

// Map of jsondb
func (d *Metadata) Map() map[string]interface{} {
	return d.m
}

// Path of jsondb
func (d *Metadata) Path() string {
	return d.path
}
