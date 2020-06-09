package main

import (
	"log"

	"{{.AppPkg}}/lib"
	"{{.AppPkg}}/lib/funcs"
  "{{.AppPkg}}/lib/hooks"
  "{{.AppPkg}}/lib/watchers"
	"github.com/guark/guark"
)

var (
	g *guark.Guark
)

func init() {

	g = guark.New()
	// g = guark.New(lib.Assets)
	// g.Watch("network", watchers.NetworkWatcher)
	// g.Watch("signale", watchers.SignalWatcher)
	//
	// Bind exported functions.
	// g.BindFuncs(lib.Funcs)

	// Bind plugins functions.
	// g.BindPlugins(lib.Plugins)

}

func main() {

	defer g.Exit()

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
