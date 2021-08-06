package middlewares

import (
	"github.com/gobuffalo/buffalo"
)

func CheckExample(next buffalo.Handler) buffalo.Handler {
	// do something
	return next
}
