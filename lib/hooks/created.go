package hooks

import (
	"github.com/guark/guark/app"
)

// App created hook.
func Created(a *app.App) error {

	a.Log.Info("---- HOOK: App created! ----")

	return nil
}
