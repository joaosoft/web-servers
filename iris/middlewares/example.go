package middlewares

import (
	"web-servers/implementation/middlewares"

	"github.com/kataras/iris/v12/context"
)

func CheckExample(ctx context.Context) {
	// do something
	_ = middlewares.ExecuteExample()
	ctx.Next()
}
