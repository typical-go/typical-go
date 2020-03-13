package typapp

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typapp/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// TypicalApp is typical application model
type TypicalApp struct {
	projectSources []string
	appModule      interface{}
	modules        []interface{}
}

// New return new instance of app
func New(appModule interface{}) *TypicalApp {
	app := &TypicalApp{
		projectSources: []string{common.PackageName(appModule)},
		appModule:      appModule,
		modules:        []interface{}{appModule},
	}
	return app
}

// WithProjectSources return app with new source
func (a *TypicalApp) WithProjectSources(sources ...string) *TypicalApp {
	a.projectSources = sources
	return a
}

// AppendModule return app with appended module. Module should be implementation of Provider, Preparer (optional) and Destroyer (optional).
func (a *TypicalApp) AppendModule(modules ...interface{}) *TypicalApp {
	a.modules = append(a.modules, modules...)
	return a
}

// AppendProjectSource return app with appended project sources
func (a *TypicalApp) AppendProjectSource(sources ...string) *TypicalApp {
	a.projectSources = append(a.projectSources, sources...)
	return a
}

// EntryPoint of app
func (a *TypicalApp) EntryPoint() *typdep.Invocation {
	if entryPointer, ok := a.appModule.(EntryPointer); ok {
		return entryPointer.EntryPoint()
	}
	return nil
}

// Provide to return constructors
func (a *TypicalApp) Provide() (constructors []*typdep.Constructor) {
	constructors = append(constructors, appConstructors...)
	for _, module := range a.modules {
		if provider, ok := module.(Provider); ok {
			constructors = append(constructors, provider.Provide()...)
		}
	}
	return
}

//Destroy to return destructor
func (a *TypicalApp) Destroy() (destructors []*typdep.Invocation) {
	for _, module := range a.modules {
		if destroyer, ok := module.(Destroyer); ok {
			destructors = append(destructors, destroyer.Destroy()...)
		}
	}
	return
}

// Prepare to return preparations
func (a *TypicalApp) Prepare() (preparations []*typdep.Invocation) {
	for _, module := range a.modules {
		if preparer, ok := module.(Preparer); ok {
			preparations = append(preparations, preparer.Prepare()...)
		}
	}
	return
}

// Commands to return commands
func (a *TypicalApp) Commands(c *Context) (cmds []*cli.Command) {
	if commander, ok := a.appModule.(Commander); ok {
		return commander.Commands(c)
	}
	return
}

// ProjectSources return source for app
func (a *TypicalApp) ProjectSources() []string {
	return a.projectSources
}

// Run application
func (a *TypicalApp) Run(d *typcore.Descriptor) (err error) {
	c := &Context{
		Descriptor: d,
		TypicalApp: a,
	}
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	if entryPoint := a.EntryPoint(); entryPoint != nil {
		app.Action = c.ActionFunc(entryPoint)
	}
	app.Commands = a.Commands(c)
	return app.Run(os.Args)
}

// Precondition the app
func (a *TypicalApp) Precondition(c *typcore.PreconditionContext) (err error) {
	var constructors []string

	if err = c.Ast().EachAnnotation("constructor", typast.FunctionType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		constructors = append(constructors, fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
		return
	}); err != nil {
		return
	}

	if c.ConfigManager != nil {
		for _, bean := range c.Configurations() {
			constructors = append(constructors, configDefinition(bean))
		}
	}

	log.Info("Generate constructors")
	target := "typical/init_app_do_not_edit.go"
	if err = a.generateConstructor(c, target, constructors); err != nil {
		return
	}
	return
}

func configDefinition(bean *typcore.Configuration) string {
	typ := reflect.TypeOf(bean.Spec()).String()
	typ2 := typ[1:]
	return fmt.Sprintf(`func(loader typcore.ConfigLoader) (cfg %s, err error){
		cfg = new(%s)
		err = loader.LoadConfig("%s", cfg)
		return 
	}`, typ, typ2, bean.Name())
}

func (a *TypicalApp) generateConstructor(c *typcore.PreconditionContext, target string, constructors []string) (err error) {
	ctx := context.Background()
	imports := []string{}
	for _, dir := range c.ProjectDirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", c.ProjectPackage, dir))
		}
	}
	if err = exor.NewWriteTemplate(target, tmpl.Constructor, tmpl.ConstructorData{
		Imports:      imports,
		Constructors: constructors,
	}).Execute(ctx); err != nil {
		return
	}
	cmd := exec.CommandContext(ctx,
		fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
