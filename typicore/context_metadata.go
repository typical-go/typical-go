package typicore

type ContextMetadata struct {
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	AppModule string   `json:"app_module"`
	Modules   []string `json:"module"`
}
