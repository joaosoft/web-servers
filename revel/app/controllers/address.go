package controllers

import (
	"fmt"
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

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	c.Response.WriteHeader(http.StatusOK, "application/json")
	return c.RenderJSON(AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	})
}
