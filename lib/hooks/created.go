package hooks

import (
	"github.com/guark/guark/app"
)

// App created hook.
func Created(app *app.App) {

	app.Log.Debug("App created hook!")

	// App created..
}
