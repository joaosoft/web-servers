package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"

	"github.com/gobuffalo/buffalo"
)

func CheckExample(next buffalo.Handler) buffalo.Handler {
	// do something
	_ = middlewares.ExecuteExample()
	return next
}
