package jsondb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	// DB is simple file-based json database
	DB struct {
		path string
		m    map[string]interface{}
	}
)

// Open path to create json database
func Open(path string) (db *DB, err error) {
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

	return &DB{
		path: path,
		m:    m,
	}, nil
}

func createEmpty(path string) (db *DB, err error) {
	db = &DB{
		path: path,
		m:    map[string]interface{}{},
	}
	err = db.Save()
	return
}

// Save db to file
func (d *DB) Save() (err error) {
	b, _ := json.Marshal(d.m)
	return ioutil.WriteFile(d.path, b, 0777)
}

// Map of jsondb
func (d *DB) Map() map[string]interface{} {
	return d.m
}

// Path of jsondb
func (d *DB) Path() string {
	return d.path
}
