package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetPersonAddressByID(ctx echo.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param("id_person"),
		IdAddress: ctx.Param("id_address"),
	}

	response := AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}

	return ctx.JSON(http.StatusOK, response)
}
