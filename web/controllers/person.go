package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"

	"github.com/joaosoft/web"
)

func GetPersonByID(ctx *web.Context) error {
	age, _ := strconv.Atoi(ctx.Request.GetParam("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Request.GetUrlParam("id_person"),
		Age:      age,
	}

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
	if err != nil {
		return ctx.Response.JSON(web.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	return ctx.Response.JSON(web.StatusOK, person)
}
