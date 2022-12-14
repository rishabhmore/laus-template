package commands

import (
	"fmt"

	e "github.com/rishabhmore/go-hustle-template/server"
	cli "github.com/urfave/cli/v2"
)

var Server *cli.Command = &cli.Command{
	Name:  "server",
	Usage: "loads the server configuration and starts the server!",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "reset",
			Aliases: []string{"r"},
			Usage:   "Will hard reset the server. Migrations will be downed and Seeders will be run",
		},
	},
	Action: serverActions,
}

func serverActions(cCtx *cli.Context) error {
	fmt.Println("Attempting to Start a Server!")

	// if cCtx.Bool("reset") {
	// 	// TODO: If reset is selected, then we have to first
	// 	// 1. Run migrations down & then up
	// 	// 2. Seed the data from seeders again
	// }

	// TODO: Load the environment variables and configuration
	// TODO: Start the echo server
	return e.StartEchoServer()
}
