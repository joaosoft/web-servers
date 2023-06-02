package middlewares

import "github.com/joaosoft/web-servers/domain/middlewares"

func CheckExample(next HandlerFunc) HandlerFunc {
	// do something
	_ = middlewares.ExecuteExample()

	return next
}
