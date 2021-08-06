package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetErrorByID(ctx *gin.Context) {
	errorID, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id_error"))
	statusText := http.StatusText(errorID)

	if statusText != "" {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, ErrorResponse{
			Code:    errorID,
			Message: statusText,
		})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
