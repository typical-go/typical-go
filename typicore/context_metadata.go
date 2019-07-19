package typicore

type ContextMetadata struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	AppModule   string   `json:"app_module"`
	Modules     []string `json:"module"`
	ProjectPath string   `json:"project_path"`
}
