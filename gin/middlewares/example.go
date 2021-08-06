package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CheckExample(ctx *gin.Context) {
	// do something
	ctx.Next()
}
