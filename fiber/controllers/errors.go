package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"
)

func GetErrorByID(ctx *fiber.Ctx) error {
	errorID, _ := strconv.Atoi(ctx.Query("id_error"))

	ctx.Response().Header.SetContentType("application/json")

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusInternalServerError)
		return ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.Response().SetStatusCode(http.StatusOK)
	return ctx.JSON(er)
}
