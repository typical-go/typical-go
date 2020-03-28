package typreadme

import (
	"fmt"
	"html/template"
	"os"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/urfave/cli/v2"
)

const (
	defaultTargetFile   = "README.md"
	defaultTemplateFile = "README.tmpl"
)

// ReadmeGenerator module
type ReadmeGenerator struct {
	// NOTE: required to be public to be access by template
	TargetFile   string
	TemplateFile string

	title       string
	description string
	usages      []UsageInfo
	buildUsages []UsageInfo
	configs     []ConfigInfo
}

// Generator retunr new instance of ReadmeGenerator
func Generator() *ReadmeGenerator {
	return &ReadmeGenerator{
		TargetFile:   defaultTargetFile,
		TemplateFile: defaultTemplateFile,
	}
}

// WithTargetFile return module with new target file
func (m *ReadmeGenerator) WithTargetFile(targetFile string) *ReadmeGenerator {
	m.TargetFile = targetFile
	return m
}

// WithTemplateFile return module with new template file
func (m *ReadmeGenerator) WithTemplateFile(templateFile string) *ReadmeGenerator {
	m.TemplateFile = templateFile
	return m
}

// WithTitle return module with new title
func (m *ReadmeGenerator) WithTitle(title string) *ReadmeGenerator {
	m.title = title
	return m
}

// WithDescription return module with new description
func (m *ReadmeGenerator) WithDescription(description string) *ReadmeGenerator {
	m.description = description
	return m
}

// WithUsages return module with new usages
func (m *ReadmeGenerator) WithUsages(usages []UsageInfo) *ReadmeGenerator {
	m.usages = usages
	return m
}

// WithBuildUsages return module with new build usages
func (m *ReadmeGenerator) WithBuildUsages(buildUsages []UsageInfo) *ReadmeGenerator {
	m.buildUsages = buildUsages
	return m
}

// WithConfigs return odule with new configs
func (m *ReadmeGenerator) WithConfigs(configs []ConfigInfo) *ReadmeGenerator {
	m.configs = configs
	return m
}

// Commands of readme
func (m *ReadmeGenerator) Commands(c *typbuildtool.Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "readme",
			Usage: "Generate README Documentation",
			Action: func(cliCtx *cli.Context) (err error) {
				return m.generate(c)
			},
		},
	}
}

func (m *ReadmeGenerator) generate(c *typbuildtool.Context) (err error) {
	var (
		file *os.File
		tmpl *template.Template
	)
	if file, err = os.Create(m.TargetFile); err != nil {
		return
	}
	defer file.Close()
	c.Infof("Parse template '%s'", m.TemplateFile)
	if tmpl, err = template.ParseFiles(m.TemplateFile); err != nil {
		return
	}
	c.Infof("Apply template and write to '%s'", m.TargetFile)
	return tmpl.Execute(file, &Object{
		TemplateFile: m.TemplateFile,
		Title:        m.Title(c),
		Description:  m.Description(c),
		Usages:       m.Usages(c),
		BuildUsages:  m.BuildUsages(c),
		Configs:      m.Configs(c),
	})
}

// Title of readme
func (m *ReadmeGenerator) Title(c *typbuildtool.Context) string {
	if m.title == "" {
		return c.Name
	}
	return m.title
}

// Description of readme
func (m *ReadmeGenerator) Description(c *typbuildtool.Context) string {
	if m.description == "" {
		return c.Description
	}
	return m.description
}

// Usages of readme
func (m *ReadmeGenerator) Usages(c *typbuildtool.Context) (infos []UsageInfo) {
	if len(m.usages) < 1 {
		if app, ok := c.App.(*typapp.App); ok {
			if app.EntryPoint() != nil {
				infos = append(infos, UsageInfo{
					Usage:       c.Name,
					Description: "Run the application",
				})
			}
			for _, cmd := range app.Commands(&typapp.Context{}) {
				infos = append(infos, usageInfos(c.Name, cmd)...)
			}
		}
		return
	}
	return m.usages
}

// BuildUsages of readme
func (m *ReadmeGenerator) BuildUsages(c *typbuildtool.Context) (infos []UsageInfo) {
	if len(m.buildUsages) < 1 {
		for _, cmd := range c.BuildTool.Commands(&typbuildtool.Context{}) {
			infos = append(infos, usageInfos("./typicalw", cmd)...)
		}
		return
	}
	return m.buildUsages
}

// Configs of readme
func (m *ReadmeGenerator) Configs(c *typbuildtool.Context) (infos []ConfigInfo) {
	if len(m.configs) < 1 {
		for _, cfg := range c.Configurations() {
			for _, field := range typcfg.RetrieveFields(cfg) {
				var required string
				if field.Required {
					required = "Yes"
				}
				infos = append(infos, ConfigInfo{
					Name:     field.Name,
					Type:     field.Type,
					Default:  field.Default,
					Required: required,
				})
			}
		}
		return
	}
	return m.configs
}

func usageInfos(name string, cmd *cli.Command) (details []UsageInfo) {
	details = append(details, UsageInfo{
		Usage:       fmt.Sprintf("%s %s", name, cmd.Name),
		Description: cmd.Usage,
	})
	for _, subcmd := range cmd.Subcommands {
		details = append(details, UsageInfo{
			Usage:       fmt.Sprintf("%s %s %s", name, cmd.Name, subcmd.Name),
			Description: subcmd.Usage,
		})
	}
	return
}
