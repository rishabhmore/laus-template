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
			{
				Name:        "generate",
				Usage:       "generate command [arguments...]",
				Description: "Command to generate dbmodels and graphql schemas",
				Aliases:     []string{"gen"},
				Subcommands: []*cli.Command{
					// TODO: add individual sub commands here
					// models
					// graphql
				},
			},
			{
				Name:        "run",
				Usage:       "run command [arguments...]",
				Description: "Command to execute project actions such as migrations, server, tests, etc.",
				Subcommands: []*cli.Command{
					// TODO: add individual sub commands here
					// migrations up/down
					// seeders
					// server --reset
					// tests
				},
			},
			// cli command to greet user
			commands.Greet,
		},
	}

	// Run the cli application! Log any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
