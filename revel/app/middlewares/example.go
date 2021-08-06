package middlewares

import (
	"fmt"

	"github.com/revel/revel"
)

var (
	CheckExample = func(c *revel.Controller, fc []revel.Filter) {
		fmt.Println("passing in the middleware example")

		fc[0](c, fc[1:])
	}
)
