package typical

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func taskPrintContext(c *typgo.BuildCli) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "desc",
			Usage: "Print descriptor",
			Action: func(cliCtx *cli.Context) (err error) {
				// b, err := json.MarshalIndent(bt.Descriptor, "", "    ")
				// b, err := json.Marshal(bt)
				// if err != nil {
				// 	return
				// }
				// fmt.Println(string(b))
				fmt.Printf("name=%s\n", c.Descriptor.Name)
				fmt.Printf("version=%s\n", c.Descriptor.Version)
				return
			},
		},
	}
}
