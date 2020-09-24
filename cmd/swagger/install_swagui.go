package swagger

import (
	"errors"

	"github.com/gookit/gcli/v2"
)

var InstallSwagUI = &gcli.Command{
	Name:    "swag:downui",
	Aliases: []string{"swag:inui", "swag:installui"},
	UseFor:  "download latest swagger-UI assets from github repository",
	Config: func(c *gcli.Command) {

	},
	Func: func(c *gcli.Command, _ []string) error {
		return errors.New("TODO")
	},
}
