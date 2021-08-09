package middlewares

import "web-servers/implementation/middlewares"

func CheckExample(next HandlerFunc) HandlerFunc {
	// do something
	_ = middlewares.ExecuteExample()

	return next
}
