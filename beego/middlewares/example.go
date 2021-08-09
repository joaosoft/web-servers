package middlewares

import (
	"web-servers/implementation/middlewares"

	"github.com/astaxie/beego/context"
)

func CheckExample(ctx *context.Context) {
	// do something
	_ = middlewares.ExecuteExample()
}
