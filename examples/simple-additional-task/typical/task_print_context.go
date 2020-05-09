package typical

import (
	"encoding/json"
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func taskPrintContext(c *typgo.Context) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "context",
			Aliases: []string{"ctx"},
			Usage:   "Print context as json",
			Action: func(cliCtx *cli.Context) (err error) {
				var b []byte
				if b, err = json.MarshalIndent(c, "", "    "); err != nil {
					return
				}
				fmt.Println(string(b))
				return
			},
		},
	}

}
