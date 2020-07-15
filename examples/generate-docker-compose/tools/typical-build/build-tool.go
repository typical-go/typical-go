package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "generate-docker-compose",
		Version: "1.0.0",
		Layouts: typgo.Layouts{"internal"},

		Commands: typgo.Commands{
			&typgo.CompileCmd{
				Action: &typgo.StdCompile{},
			},
			&typgo.RunCmd{
				Action: &typgo.StdRun{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},
			&typdocker.Command{
				Composers: []typdocker.Composer{
					redisRecipe,
				},
			},
		},
	}

	redisRecipe = &typdocker.Recipe{
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
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
