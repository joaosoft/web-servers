package controllers

import (
	"encoding/json"
	"net/http"
	"web-servers/implementation/models"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetErrorByID(ctx *routing.Context) error {
	errorID := ctx.QueryArgs().GetUintOrZero("id_error")

	ctx.SetContentType("application/json")

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return ctx.WriteData(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	bytes, _ := json.Marshal(er)

	ctx.SetStatusCode(http.StatusOK)
	_, err = ctx.Write(bytes)
	return err
}
