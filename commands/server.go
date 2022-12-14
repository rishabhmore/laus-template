package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var Server *cli.Command = &cli.Command{
	Name:   "server",
	Usage:  "loads the server configuration and starts the server!",
	Action: serverActions,
}

func serverActions(*cli.Context) error {
	fmt.Println("Attempting to Start a Server!")

	// TODO: Load the environment variables and configuration
	// TODO: Start the echo server
	return nil
}
