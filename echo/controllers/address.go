package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func GetPersonAddressByID(ctx echo.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param(":id_person"),
		IdAddress: ctx.Param(":id_address"),
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	response :=	AddressResponse{
			Id:      request.IdAddress,
			Country: "Portugal",
			City:    "Porto",
			Street:  "Rua da calçada",
			Number:  7,
		}

	return ctx.JSON(http.StatusOK, response)
}
