package controllers

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func GetPersonAddressByID(ctx iris.Context) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Params().Get("id_person"),
		IdAddress: ctx.Params().Get("id_address"),
	}

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
