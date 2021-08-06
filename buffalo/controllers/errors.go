package controllers

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"net/http"
	"strconv"
)

func GetErrorByID(ctx buffalo.Context) error {
	errorID, _ := strconv.Atoi(ctx.Request().URL.Query().Get("id_error"))
	fmt.Printf("> executing get errors for id: %d", errorID)

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
