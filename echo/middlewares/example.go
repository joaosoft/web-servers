package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"

	"github.com/labstack/echo"
)

func CheckExample(next echo.HandlerFunc) echo.HandlerFunc {
	// do something
	_ = middlewares.ExecuteExample()
	return next
}
