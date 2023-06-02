package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"
)

func GetErrorByID(ctx iris.Context) {
	errorID, _ := strconv.Atoi(ctx.URLParam("id_error"))

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		_, _ = ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(er)
}
