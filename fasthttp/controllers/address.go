package controllers

import (
	"encoding/json"
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetPersonAddressByID(ctx *routing.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param("id_person"),
		IdAddress: ctx.Param("id_address"),
	}

	ctx.SetContentType("application/json")

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return ctx.WriteData(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	bytes, _ := json.Marshal(address)

	ctx.SetStatusCode(http.StatusOK)
	_, err = ctx.Write(bytes)
	return err
}
