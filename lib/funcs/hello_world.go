package funcs

import (
	"github.com/guark/guark/app"
)

func HelloWorld(c app.Context) (interface{}, error) {

	c.App.Log.Info("HelloWorld called")

	return nil, nil
}
