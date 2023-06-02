package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetPersonByID(ctx echo.Context) error {
	age, _ := strconv.Atoi(ctx.QueryParam("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
		Age:      age,
	}

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, age)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return ctx.JSON(http.StatusOK, person)
}
