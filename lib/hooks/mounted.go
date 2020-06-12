package hooks

import (
	"github.com/guark/guark/app"
)

// App mounted hook.
func Mounted(app *app.App) {

	app.Log.Debug("---- HOOK: App mounted! ----")

	// App created..
}
