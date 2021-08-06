package middlewares

import (
	"github.com/labstack/echo"
)

func CheckExample(next echo.HandlerFunc) echo.HandlerFunc {
	// do something
	return next
}
