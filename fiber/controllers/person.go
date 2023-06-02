package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"
)

func GetPersonByID(ctx *fiber.Ctx) error {
	age, _ := strconv.Atoi(ctx.Query("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Params("id_person"),
		Age:      age,
	}

	ctx.Response().Header.SetContentType("application/json")

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusInternalServerError)
		return ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.Response().SetStatusCode(http.StatusOK)
	return ctx.JSON(person)
}
