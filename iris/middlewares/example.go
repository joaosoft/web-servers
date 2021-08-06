package middlewares

import (
	"fmt"

	"github.com/kataras/iris/v12/context"
)

func CheckExample(ctx context.Context) {
	fmt.Println("passing in the middleware example")
	ctx.Next()
}
