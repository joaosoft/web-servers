package controllers

import (
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/gofiber/fiber"
)

func GetPersonByID(ctx *fiber.Ctx) {
	age, _ := strconv.Atoi(ctx.Query("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Params("id_person"),
		Age:      age,
	}

	ctx.Set("Content-Type", "application/json")

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_ = ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.Status(http.StatusOK)
	_ = ctx.JSON(person)
}
