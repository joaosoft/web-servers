package controllers

import (
	"net/http"
	"web-servers/implementation/models"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetPersonAddressByID(ctx *routing.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param("id_person"),
		IdAddress: ctx.Param("id_address"),
	}

	ctx.SetContentType("application/json")

	person, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return ctx.WriteData(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.SetStatusCode(http.StatusOK)
	return ctx.WriteData(person)
}
