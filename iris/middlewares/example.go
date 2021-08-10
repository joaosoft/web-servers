package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/kataras/iris/context"
)

func CheckExample(ctx context.Context) {
	// do something
	_ = middlewares.ExecuteExample()
	ctx.Next()
}
