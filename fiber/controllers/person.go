package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/gofiber/fiber/v2"
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

	bytes, _ := json.Marshal(person)

	ctx.Response().SetStatusCode(http.StatusOK)
	_, err = ctx.Write(bytes)
	return err
}
