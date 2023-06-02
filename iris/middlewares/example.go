package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"

	"github.com/kataras/iris/v12/context"
)

func CheckExample(ctx context.Context) {
	// do something
	_ = middlewares.ExecuteExample()
	ctx.Next()
}
