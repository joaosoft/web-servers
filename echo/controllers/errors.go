package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetErrorByID(ctx echo.Context) error {
	errorID, _ := strconv.Atoi(ctx.QueryParam("id_error"))

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return ctx.JSON(http.StatusOK, er)
}
