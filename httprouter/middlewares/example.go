package middlewares

import (
	"fmt"
)

func CheckExample(next HandlerFunc) HandlerFunc {
	fmt.Println("passing in the middleware example")
	return next
}
