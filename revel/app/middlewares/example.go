package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/revel/revel"
)

var (
	CheckExample = func(c *revel.Controller, fc []revel.Filter) {
		// do something
		_ = middlewares.ExecuteExample()
		fc[0](c, fc[1:])
	}
)
