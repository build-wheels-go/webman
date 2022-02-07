package main

import "webman/framework"

func registerRouter(core *framework.Core) {
	core.Set("foo", FooControllerHandler)
}
