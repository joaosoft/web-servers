package middlewares

import (
	"net/http"
)

func CheckExample(next http.Handler) http.Handler {
	// do something
	return next
}
