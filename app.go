package main

import (
	"log"

	"github.com/guark/guark"
	"github.com/guark/guark/app"
	"{{AppPkg}}/lib"
)

var (
	g *guark.Guark
)

func init() {

	g = guark.New(&app.Config{
		Hooks:    lib.Hooks,
		Funcs:    lib.Funcs,
		Assets:   lib.Assets,
		Plugins:  lib.Plugins,
		Watchers: lib.Watchers,
	})
}

func main() {

	defer g.Exit()

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
