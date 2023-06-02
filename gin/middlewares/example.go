package middlewares

import (
	"github.com/joaosoft/web-servers/domain/middlewares"

	"github.com/gin-gonic/gin"
)

func CheckExample(ctx *gin.Context) {
	// do something
	_ = middlewares.ExecuteExample()
	ctx.Next()
}
