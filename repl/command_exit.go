package repl

import (
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
	"os"
)

func CallbackExit(cfg *config.Config, args ...string) error {
	os.Exit(0)
	return nil
}
