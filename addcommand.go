package cli

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

// AddCommand adds a new command to the application. The command must have a
// special field defining name, short-description and long-description (see
// package documentation). It panics if the command is not valid.
// Returned CommandAdder can be used to add subcommands.
//
// Additional functions can be passed to manipulate the resulting *flags.Command
// after its initialization.
func (a *App) AddCommand(cmd interface{}, cfs ...func(*flags.Command)) CommandAdder {
	return commandAdder{a.Parser}.AddCommand(cmd, cfs...)
}

// CommandAdder can be used to add subcommands.
type CommandAdder interface {
	// AddCommand adds the given commands as subcommands of the
	// cuurent one.
	AddCommand(interface{}, ...func(*flags.Command)) CommandAdder
}

type commandAdder struct {
	internalCommandAdder
}

type internalCommandAdder interface {
	AddCommand(string, string, string, interface{}) (*flags.Command, error)
}

func (a commandAdder) AddCommand(cmd interface{}, cfs ...func(*flags.Command)) CommandAdder {
	typ, err := getStructType(cmd)
	if err != nil {
		panic(err)
	}

	under, ok := typ.FieldByName("_")
	if !ok {
		panic(fmt.Errorf("missing `_` field"))
	}

	name := under.Tag.Get("name")
	shortDescription := under.Tag.Get("short-description")
	longDescription := under.Tag.Get("long-description")

	if v, ok := cmd.(ContextCommander); ok {
		cmd = &nopCommander{v}
	}

	c, err := a.internalCommandAdder.AddCommand(
		name, shortDescription, longDescription, cmd)
	if err != nil {
		panic(err)
	}

	for _, cf := range cfs {
		cf(c)
	}

	return commandAdder{c}
}
