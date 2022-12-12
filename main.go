package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	// Initialize a cli application
	app := &cli.App{
		Commands: []*cli.Command{
			// Create a new cli command to greet user
			{
				Name:  "greet",
				Usage: "fight the loneliness!",
				Action: func(*cli.Context) error {
					fmt.Println("Hello friend!")
					return nil
				},
			},
		},
	}

	// Run the cli application! Log any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
