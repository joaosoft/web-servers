package middlewares

import (
	"github.com/kataras/iris/v12/context"
)

func CheckExample(ctx context.Context) {
	// do something
	ctx.Next()
}
