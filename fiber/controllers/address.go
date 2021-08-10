package controllers

import (
	"net/http"
	"web-servers/domain/models"

	"github.com/gofiber/fiber"
)

func GetPersonAddressByID(ctx *fiber.Ctx) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Params("id_person"),
		IdAddress: ctx.Params("id_address"),
	}

	ctx.Set("Content-Type", "application/json")

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_ = ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.Status(http.StatusOK)
	_ = ctx.JSON(address)
}
