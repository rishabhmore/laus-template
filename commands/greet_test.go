package commands

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	cli "github.com/urfave/cli/v2"
)

func Test_executeGreet(t *testing.T) {
	tests := map[string]struct {
		args func() []string
	}{
		"Greet Friend": {
			args: func() []string {
				args := os.Args[0:1]
				args = append(args, "greet")
				return args
			},
		},
		"Greet Foe": {
			args: func() []string {
				args := os.Args[0:1:2]
				args = append(args, "greet")
				args = append(args, "-foe")
				return args
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// When we create a test instance of our app, we need to register our greet command to it
			// otherwise, the app won't recognize our command and it's flags
			app := &cli.App{
				Commands: []*cli.Command{
					// cli command to greet user
					Greet,
				},
			}
			err := app.Run(tt.args())

			if name == "Greet Friend" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
