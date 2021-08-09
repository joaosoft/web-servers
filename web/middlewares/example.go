package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/joaosoft/web"
)

func CheckExample(next web.HandlerFunc) web.HandlerFunc {
	// do something
	_ = middlewares.ExecuteExample()
	return next
}
