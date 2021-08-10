package controllers

import (
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/kataras/iris"
)

func GetErrorByID(ctx iris.Context) {
	errorID, _ := strconv.Atoi(ctx.URLParam("id_error"))

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(er)
}
