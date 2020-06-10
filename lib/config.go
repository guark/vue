package lib

import (
	"github.com/guark/guark/app"
	"{{AppPkg}}/lib/funcs"
	"{{AppPkg}}/lib/hooks"
)

// Exposed functions to guark Javascript api.
var Funcs = app.Funcs{
	"hello_my_func": funcs.FuncName,
}

// App hooks.
var Hooks = app.Hooks{
	"created": hooks.Created,
	"mounted": hooks.Mounted,
}

// App plugins.
var Plugins = app.Plugins{}

// App watchers.
var Watchers = []app.Watcher{}
