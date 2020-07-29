package main

import (
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	Name:    "docker-command",
	Version: "1.0.0",
	Layouts: []string{"internal"},

	Cmds: []typgo.Cmd{
		// compile
		&typgo.CompileCmd{
			Action: &typgo.StdCompile{},
		},

		// run
		&typgo.RunCmd{
			Before: typgo.BuildSysRuns{"compile"},
			Action: &typgo.StdRun{},
		},

		// clean
		&typgo.CleanCmd{
			Action: &typgo.StdClean{},
		},

		// docker
		&typdocker.DockerCmd{
			Composers: []typdocker.Composer{
				&typdocker.Recipe{
					Services: typdocker.Services{
						"redis": typdocker.Service{
							Image: "redis:4.0.5-alpine",
							Ports: []string{"6379:6379"},
						},
						"webdis": typdocker.Service{
							Image: "anapsix/webdis",
							Ports: []string{"7379:7379"},
						},
					},
				},
			},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
