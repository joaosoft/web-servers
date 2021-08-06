package controllers

import (
	"encoding/json"
	"net/http"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetPersonAddressByID(ctx *routing.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param("id_person"),
		IdAddress: ctx.Param("id_address"),
	}

	ctx.SetContentType("application/json")

	bytes, err := json.Marshal(
		AddressResponse{
			Id:      request.IdAddress,
			Country: "Portugal",
			City:    "Porto",
			Street:  "Rua da calçada",
			Number:  7,
		},
	)

	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return nil
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.Write(bytes)

	return nil
}
