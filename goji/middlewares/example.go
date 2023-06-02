package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"
	"net/http"
)

func CheckExample(next http.Handler) http.Handler {
	// do something
	_ = middlewares.ExecuteExample()
	return next
}
