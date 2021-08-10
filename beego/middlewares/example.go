package middlewares

import (
	"web-servers/domain/middlewares"

	"github.com/astaxie/beego/context"
)

func CheckExample(ctx *context.Context) {
	// do something
	_ = middlewares.ExecuteExample()
}
