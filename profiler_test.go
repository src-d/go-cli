package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type NopCommand struct {
	_ string `name:"nop" short-description:"nop" long-description:"nop"`
	Command
}

func (c *NopCommand) Execute(args []string) error {
	return nil
}

func setupDefaultCommand(t *testing.T) *App {
	app := New("test", "", "", "")
	app.AddCommand(&NopCommand{})
	return app
}

func TestProfilerOptions_Enable(t *testing.T) {
	require := require.New(t)
	app := setupDefaultCommand(t)
	err := app.Run([]string{"test", "nop", "--profiler-http", "--profiler-block-rate", "10"})
	require.NoError(err)
}
func TestProfilerOptions_Error(t *testing.T) {
	require := require.New(t)
	app := setupDefaultCommand(t)
	err := app.Run([]string{"test", "nop", "--profiler-http", "--profiler-endpoint", "a.b.c.d:foo"})
	require.Error(err)
}
