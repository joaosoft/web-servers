package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/go-martini/martini"
)

func CheckExample(c martini.Context) {
	// do something
	_ = middlewares.ExecuteExample()
	c.Next()
}
