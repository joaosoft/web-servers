package middlewares

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
)

func CheckExample(next buffalo.Handler) buffalo.Handler {
	fmt.Println("passing in the middleware example")
	return next
}
