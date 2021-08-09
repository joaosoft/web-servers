package middlewares

import (
	"net/http"
	"web-servers/implementation/middlewares"
)

func CheckExample(next http.Handler) http.Handler {
	// do something
	_ = middlewares.ExecuteExample()
	return next
}
