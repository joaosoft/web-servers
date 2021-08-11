package controllers

import (
	"net/http"
	"web-servers/domain/models"

	"github.com/gofiber/fiber/v2"
)

func GetPersonAddressByID(ctx *fiber.Ctx) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Params("id_person"),
		IdAddress: ctx.Params("id_address"),
	}

	ctx.Response().Header.SetContentType("application/json")

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusInternalServerError)
		return ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.Response().SetStatusCode(http.StatusOK)
	return ctx.JSON(address)
}
