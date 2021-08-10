package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/gofiber/fiber"
)

func CheckExample(ctx *fiber.Ctx) {
	// do something
	_ = middlewares.ExecuteExample()
	ctx.Next()
}
