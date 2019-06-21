package appx

import "gopkg.in/urfave/cli.v1"

// DBInfra database infrastructure
type DBInfra interface {
	Create() (err error)
	Drop() (err error)
	Migrate(args cli.Args) error
	Rollback(args cli.Args) error
}
