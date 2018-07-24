package main

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/src-d/go-cli.v0"
)

func init() {
	app.AddCommand(&sleepCommand{})
}

type sleepCommand struct {
	cli.Command `name:"sleep" short-description:"sleeps forever" long-description:"sleeps indefinitely until it receives SIGTERM or SIGINT"`

	Positional struct {
		Sleep time.Duration `positional-arg-name:"sleep" default:"1s" description:"sleep intervals"`
	} `positional-args:"yes"`
}

func (c *sleepCommand) ExecuteContext(ctx context.Context, args []string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		fmt.Println("Sleeping...")
		time.Sleep(c.Positional.Sleep)
	}
}
