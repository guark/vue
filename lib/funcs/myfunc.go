package funcs

import (
	"fmt"

	"github.com/guark/guark/app"
)

func FuncName(c app.Context) (interface{}, error) {

	fmt.Println("Func called " + c.Params.Get("name").(string))

	return nil, nil
}
