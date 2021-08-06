package controllers

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"net/http"
	"strconv"
)

func GetErrorByID(ctx buffalo.Context) error {
	errorID, _ := strconv.Atoi(ctx.Request().URL.Query().Get("id_error"))
	statusText := http.StatusText(errorID)

	if statusText != "" {
		response := ErrorResponse{
				Code:    errorID,
				Message: statusText,
			}
		return ctx.Render(http.StatusOK, render.JSON(response))
	} else {
		return ctx.Render(http.StatusNoContent, nil)
	}
}
