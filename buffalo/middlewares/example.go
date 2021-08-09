package middlewares

import (
	"web-servers/implementation/middlewares"

	"github.com/gobuffalo/buffalo"
)

func CheckExample(next buffalo.Handler) buffalo.Handler {
	// do something
	_ = middlewares.ExecuteExample()
	return next
}
