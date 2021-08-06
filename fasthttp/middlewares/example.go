package middlewares

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func CheckExample(ctx *routing.Context) error {
	fmt.Println("passing in the middleware example")
	return ctx.Next()
}
