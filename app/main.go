package main

import (
	"os"

	"app/common/config"
	"app/web"

	"github.com/urfave/cli"
)

var cfg = config.GetConfig()

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "web",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				return web.Tokoin.Start()
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	app.Run(os.Args)
}
