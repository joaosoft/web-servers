package middlewares

import (
	"fmt"
)

func CheckExample(HandlerFunc) HandlerFunc {
	fmt.Println("passing in the middleware example")
	return nil
}
