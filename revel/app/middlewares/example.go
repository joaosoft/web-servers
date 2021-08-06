package middlewares

import (
	"github.com/revel/revel"
)

var (
	CheckExample = func(c *revel.Controller, fc []revel.Filter) {
		// do something

		fc[0](c, fc[1:])
	}
)
