package typreadme

// Data to supply to readme template
type Data struct {
	TemplateFile string
	Title        string
	Description  string
	Usages       []UsageInfo
	BuildUsages  []UsageInfo
	Configs      []ConfigInfo
}

// UsageInfo is command information
type UsageInfo struct {
	Usage       string
	Description string
}

// ConfigInfo is configuration information
type ConfigInfo struct {
	Name     string
	Type     string
	Default  string
	Required string
}
