package middlewares

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

func CheckExample(ctx *context.Context) {
	fmt.Println("passing in the middleware example")
}
