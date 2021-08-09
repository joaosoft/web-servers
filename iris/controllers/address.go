package controllers

import (
	"net/http"
	"web-servers/domain/models"

	"github.com/kataras/iris/v12"
)

func GetPersonAddressByID(ctx iris.Context) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Params().Get("id_person"),
		IdAddress: ctx.Params().Get("id_address"),
	}

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
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
	ctx.JSON(address)
}
