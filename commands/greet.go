package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var Greet *cli.Command = &cli.Command{
	Name:  "greet",
	Usage: "fight the loneliness!",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "foe",
			Aliases: []string{"f"},
			Usage:   "You dare to greet your foe?",
		},
	},
	Action: executeGreet,
}

func executeGreet(cliCtx *cli.Context) error {
	if cliCtx.Bool("foe") {
		return fmt.Errorf("foes are not welcome here")
	} else {
		fmt.Println("Hello friend!")
		return nil
	}
}
