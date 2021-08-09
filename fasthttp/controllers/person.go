package controllers

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"net/http"
	models2 "web-servers/implementation/models"
)

func GetPersonByID(ctx *routing.Context) error {
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
		Age:      ctx.QueryArgs().GetUintOrZero("age"),
	}

	ctx.SetContentType("application/json")

	person, err := (&models2.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return ctx.WriteData(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.SetStatusCode(http.StatusOK)
	return ctx.WriteData(person)
}
