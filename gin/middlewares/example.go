package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckExample(ctx *gin.Context) {
	fmt.Println("passing in the middleware example")
	ctx.Next()
}
