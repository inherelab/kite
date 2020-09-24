package cmd

import (
	"errors"

	"github.com/gookit/gcli/v2"
)

var HotReloadServe = &gcli.Command{
	Name:    "hotreload",
	Aliases: []string{"hot:reload"},
	UseFor:  "hot reload serve on files modified",
	Config: func(c *gcli.Command) {

	},
	Func: func(c *gcli.Command, _ []string) error {
		return errors.New("TODO")
	},
}
