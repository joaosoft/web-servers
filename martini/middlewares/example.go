package middlewares

import (
	"fmt"
	"github.com/go-martini/martini"
)

func CheckExample(c martini.Context) {
	fmt.Println("passing in the middleware example")
	c.Next()
}
