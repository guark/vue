package hooks

import (
	"github.com/guark/guark/app"
)

// App created hook.
func Created(app *app.App) {

	app.Log.Debug("---- HOOK: App created! ----")

	// App created..
}
