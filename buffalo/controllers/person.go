package controllers

import (
	"net/http"
	"strconv"
	"web-servers/implementation/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func GetPersonByID(ctx buffalo.Context) error {
	age, _ := strconv.Atoi(ctx.Request().URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
		Age:      age,
	}

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, age)
	if err != nil {
		return ctx.Render(http.StatusInternalServerError, render.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
	}

	return ctx.Render(http.StatusOK, render.JSON(person))
}
