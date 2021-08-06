package controllers

import (
	"net/http"

	"github.com/joaosoft/web"
)

func GetPersonAddressByID(ctx *web.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Request.GetUrlParam("id_person"),
		IdAddress: ctx.Request.GetUrlParam("id_address"),
	}

	response := AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}

	return ctx.Response.JSON(http.StatusOK, response)
}
