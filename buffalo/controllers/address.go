package controllers

import (
	"net/http"
	"web-servers/implementation/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func GetPersonAddressByID(ctx buffalo.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param("id_person"),
		IdAddress: ctx.Param("id_address"),
	}

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		return ctx.Render(http.StatusInternalServerError, render.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
	}

	return ctx.Render(http.StatusOK, render.JSON(address))
}
