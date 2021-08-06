package controllers

import (
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
	return c.RenderJSON(AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	})
}
