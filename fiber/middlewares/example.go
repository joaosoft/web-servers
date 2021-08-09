package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CheckExample(ctx *fiber.Ctx) error {
	// do something
	_ = middlewares.ExecuteExample()
	return ctx.Next()
}
