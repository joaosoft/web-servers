package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"

	"github.com/revel/revel"
)

type AddressController struct {
	*revel.Controller
}

func (c AddressController) GetPersonAddressByID() revel.Result {
	request := GetPersonAddressByIDRequest{
		IdPerson:  c.Params.Get("id_person"),
		IdAddress: c.Params.Get("id_address"),
	}

	c.Response.WriteHeader(http.StatusOK, "application/json")

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		c.RenderJSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	return c.RenderJSON(address)
}
