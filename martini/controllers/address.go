package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func GetPersonAddressByID(params martini.Params, r render.Render) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  params["id_person"],
		IdAddress: params["id_address"],
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	r.JSON(http.StatusOK, AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	})
}
