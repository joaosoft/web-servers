package controllers

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris"
)

func GetPersonAddressByID(ctx iris.Context) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Params().Get("id_person"),
		IdAddress: ctx.Params().Get("id_address"),
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	response := AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(response)
}
