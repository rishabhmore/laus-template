package main_test

import (
	"os"
	"testing"

	main "github.com/rishabhmore/laus-template"
	"github.com/stretchr/testify/assert"
)

func TestRunCliApp(t *testing.T) {
	tests := map[string]struct {
		args func() []string
	}{
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

			err := main.RunCliApp(tt.args())
			assert.Error(t, err)
		})
	}
}
