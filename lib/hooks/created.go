package hooks

import (
	"github.com/guark/guark/app"
)

type Created struct {
}

// App created hook.
func (h Created) Handle(app *app.App) error {

	return nil
}
