package controllers

import (
	"net/http"
	"web-servers/implementation/models"

	"github.com/joaosoft/web"
)

func GetPersonAddressByID(ctx *web.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Request.GetUrlParam("id_person"),
		IdAddress: ctx.Request.GetUrlParam("id_address"),
	}

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	return ctx.Response.JSON(web.StatusOK, address)
}
