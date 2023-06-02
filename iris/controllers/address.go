package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
)

func GetPersonAddressByID(ctx iris.Context) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Params().Get("id_person"),
		IdAddress: ctx.Params().Get("id_address"),
	}

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
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
	_, _ = ctx.JSON(address)
}
