package main

import "github.com/leijiangnan/nanhade/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
