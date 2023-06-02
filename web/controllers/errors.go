package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"

	"github.com/joaosoft/web"
)

func GetErrorByID(ctx *web.Context) error {
	errorID, _ := strconv.Atoi(ctx.Request.GetParam("id_error"))

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	return ctx.Response.JSON(web.StatusOK, er)
}
