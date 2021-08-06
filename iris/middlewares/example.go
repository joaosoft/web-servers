package middlewares

import (
	"fmt"

	"github.com/kataras/iris/context"
)

func CheckExample(ctx context.Context) {
	fmt.Println("passing in the middleware example")
	ctx.Next()
}
