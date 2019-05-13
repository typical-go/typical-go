package metadata

type Context struct {
	Name         string       `json:"name"`
	Path         string       `json:"path"`
	Description  string       `json:"path"`
	Architecture Architecture `json:"architecture"`
}
