package controllers

import (
	"net/http"
	"web-servers/implementation/models"

	"github.com/labstack/echo"
)

func GetPersonAddressByID(ctx echo.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param("id_person"),
		IdAddress: ctx.Param("id_address"),
	}

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return ctx.JSON(http.StatusOK, address)
}
