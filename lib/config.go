package lib

import (
	"github.com/guark/guark/app"
	"github.com/guark/plugins/dialog"
	"github.com/guark/plugins/notify"
	"github.com/guark/plugins/clipboard"
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
var Plugins = app.Plugins{
	"dialog": &dialog.Plugin{},
	"notify": &notify.Plugin{},
	"clipboard": &clipboard.Plugin{},
}

// App watchers.
var Watchers = []app.Watcher{}
