package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetErrorByID(ctx echo.Context) error {
	errorID, _ := strconv.Atoi(ctx.QueryParam("id_error"))
	statusText := http.StatusText(errorID)

	if statusText != "" {
		response := ErrorResponse{
			Code:    errorID,
			Message: statusText,
		}
		return ctx.JSON(http.StatusOK, response)
	} else {
		return ctx.NoContent(http.StatusNoContent)
	}
}
