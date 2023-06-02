package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"

	routing "github.com/qiangxue/fasthttp-routing"
)

func CheckExample(ctx *routing.Context) error {
	// do something
	_ = middlewares.ExecuteExample()
	return ctx.Next()
}
