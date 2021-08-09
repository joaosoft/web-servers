package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/gocraft/web"
)

func CheckExample(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	// do something
	_ = middlewares.ExecuteExample()
	next(rw, req)
}
