package middlewares

import (
	"github.com/gocraft/web"
)

func CheckExample(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	// do something
	next(rw, req)
}
