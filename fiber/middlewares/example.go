package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"
)

func CheckExample(ctx *fiber.Ctx) error {
	// do something
	_ = middlewares.ExecuteExample()
	return ctx.Next()
}
