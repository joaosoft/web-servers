package middlewares

import (
	"fmt"
	"github.com/gocraft/web"
)

func CheckExample(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	fmt.Println("passing in the middleware example")
	next(rw, req)
}
