package middlewares

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

func CheckExample(ctx *routing.Context) error {
	// do something
	return ctx.Next()
}
