package typicore

type ContextMetadata struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	ArcheType   string   `json:"archetype"`
	Modules     []string `json:"module"`
	PackageName string   `json:"package_name"`
	ProjectPath string   `json:"project_path"`
}
