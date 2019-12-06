package typprebuilder

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/typprebuilder/metadata"
	"github.com/typical-go/typical-go/pkg/utility/filekit"
)

const (
	debugEnv = "PREBUILDER_DEBUG"
)

// Run the prebuilder
func Run(ctx *typctx.Context) {
	var err error
	var preb prebuilder
	os.Mkdir(typenv.Layout.Metadata, 0700)
	os.Mkdir(typenv.DependencyPath, 0700)
	checker := checker{
		Context:         ctx,
		contextChecksum: contextChecksum(),
		buildToolBinary: !filekit.IsExist(typenv.BuildToolBin),
		readmeFile:      !filekit.IsExist(typenv.Readme),
	}
	if os.Getenv(debugEnv) != "" {
		log.SetLevel(log.DebugLevel)
	}
	if err = ctx.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	if err = GenerateEnvfile(ctx); err != nil {
		log.Fatal(err.Error())
	}
	if err := preb.Initiate(ctx); err != nil {
		log.Fatal(err.Error())
	}
	if checker.configuration, err = metadata.Update("config_fields", preb.ConfigFields); err != nil {
		log.Fatal(err.Error())
	}
	if checker.buildCommands, err = metadata.Update("build_commands", preb.BuildCommands); err != nil {
		log.Fatal(err.Error())
	}
	if checker.testTarget, err = Generate("test_target", testTarget{ContextImport: preb.ContextImport, Packages: preb.Dirs}); err != nil {
		log.Fatal(err.Error())
	}
	if checker.mockTarget, err = Generate("mock_target", mockTarget{ApplicationImports: preb.ApplicationImports, MockTargets: preb.ProjectFiles.Automocks()}); err != nil {
		log.Fatal(err.Error())
	}
	if _, err = Generate("constructor", constructor{ApplicationImports: preb.ApplicationImports, Constructors: preb.ProjectFiles.Autowires()}); err != nil {
		log.Fatal(err.Error())
	}
	if checker.checkBuildTool() {
		log.Info("Build the build-tool")
		cmd := exec.Command("go", "build",
			"-o", typenv.BuildToolBin,
			"./"+typenv.BuildToolMainPath,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err.Error())
		}
	}
	if checker.checkReadme() {
		log.Info("Generate readme")
		cmd := exec.Command(typenv.BuildToolBin, "readme")
		if err := cmd.Run(); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func contextChecksum() bool {
	// NOTE: context checksum is passed by typicalw
	if len(os.Args) > 1 {
		return os.Args[1] == "1"
	}
	return false
}

func goimports(filename string) error {
	cmd := exec.Command(fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", filename)
	return cmd.Run()
}
