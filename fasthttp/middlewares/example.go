package middlewares

import (
	"web-servers/implementation/middlewares"

	routing "github.com/qiangxue/fasthttp-routing"
)

func CheckExample(ctx *routing.Context) error {
	// do something
	_ = middlewares.ExecuteExample()
	return ctx.Next()
}
