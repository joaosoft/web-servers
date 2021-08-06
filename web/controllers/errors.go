package controllers

import (
	"net/http"
	"strconv"

	"github.com/joaosoft/web"
)

func GetErrorByID(ctx *web.Context) error {
	errorID, _ := strconv.Atoi(ctx.Request.GetParam("id_error"))
	statusText := http.StatusText(errorID)

	if statusText != "" {
		response := ErrorResponse{
			Code:    errorID,
			Message: statusText,
		}
		return ctx.Response.JSON(http.StatusOK, response)
	} else {
		return ctx.Response.NoContent(http.StatusNoContent)
	}
}
