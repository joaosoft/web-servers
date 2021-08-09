package controllers

import (
	"encoding/json"
	"net/http"
	"web-servers/implementation/models"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetPersonByID(ctx *routing.Context) error {
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
		Age:      ctx.QueryArgs().GetUintOrZero("age"),
	}

	ctx.SetContentType("application/json")

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return ctx.WriteData(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	bytes, _ := json.Marshal(person)

	ctx.SetStatusCode(http.StatusOK)
	_, err = ctx.Write(bytes)
	return err
}
