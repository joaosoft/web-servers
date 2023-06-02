package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"

	"github.com/astaxie/beego/context"
)

func CheckExample(ctx *context.Context) {
	// do something
	_ = middlewares.ExecuteExample()
}
