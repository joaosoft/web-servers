package middlewares

import (
	"fmt"
	"net/http"
)

func CheckExample(next http.Handler) http.Handler {
	fmt.Println("passing in the middleware example")
	return next
}
