package main

import (
	"log"
	"os"

	"github.com/rishabhmore/go-hustle-template/commands"
	cli "github.com/urfave/cli/v2"
)

func main() {
	// Initialize a cli application
	app := &cli.App{
		Commands: []*cli.Command{
			// Create a new cli command to greet user
			commands.Greet,
		},
	}

	// Run the cli application! Log any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
