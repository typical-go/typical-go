package typgo

import (
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

type (
	// CliContext is context of build
	CliContext struct {
		typlog.Logger

		Cli *cli.Context
		*BuildTool
	}

	// CliFunc is command line function
	CliFunc func(*CliContext) error
)
