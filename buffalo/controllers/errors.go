package controllers

import (
	"net/http"
	"strconv"
	"web-servers/implementation/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func GetErrorByID(ctx buffalo.Context) error {
	errorID, _ := strconv.Atoi(ctx.Request().URL.Query().Get("id_error"))

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		return ctx.Render(http.StatusInternalServerError, render.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
	}

	return ctx.Render(http.StatusOK, render.JSON(er))
}
