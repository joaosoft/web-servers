package middlewares

import (
	"fmt"

	"github.com/labstack/echo"
)

func CheckExample(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("passing in the middleware example")
	return next
}
