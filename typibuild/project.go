package typibuild

type Project struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	ArcheType   ArcheType `json:"archetype"`
	Modules     []string  `json:"module"`
	PackageName string    `json:"package_name"`
	Path        string    `json:"path"`
}
