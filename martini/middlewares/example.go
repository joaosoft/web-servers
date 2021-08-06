package middlewares

import (
	"github.com/go-martini/martini"
)

func CheckExample(c martini.Context) {
	// do something
	c.Next()
}
