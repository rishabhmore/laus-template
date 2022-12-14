package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var Greet *cli.Command = &cli.Command{
	Name:  "greet",
	Usage: "fight the loneliness!",
	Action: func(*cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	},
}
