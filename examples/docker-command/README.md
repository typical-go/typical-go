# Generate Docker Compose

Example typical-go project to demonstrate how to generate docker-compose

Create docker-compose recipe
```go
var redisRecipe = &typdocker.Recipe{
	Version: typdocker.V3,
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
```

Register the docker utility
```go
var Descriptor = typgo.Descriptor{
	// ...

	Utility: typdocker.Compose(
		redisRecipe,
		// More recipe...
	),
}
```

`typicalw docker` to see docker utility
```bash
./typicalw docker
```
```
NAME:
   generate-docker-compose docker - Docker utility

USAGE:
   generate-docker-compose docker command [command options] [arguments...]

COMMANDS:
   compose  Generate docker-compose.yaml
   up       Spin up docker containers according docker-compose
   down     Take down all docker containers according docker-compose
   wipe     Kill all running docker container
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

Spin up the docker
```bash
./typicalw docker compose # generate docker-compose.yml (if required)
./typicalw docker up
```