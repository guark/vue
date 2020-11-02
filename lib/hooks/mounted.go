package hooks

import (
	"github.com/guark/guark/app"
)

// App mounted hook.
func Mounted(a *app.App) error {

	a.Log.Info("---- HOOK: App mounted! ----")

	return nil
}
