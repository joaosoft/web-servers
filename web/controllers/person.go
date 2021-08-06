package controllers

import (
	"strconv"

	"github.com/joaosoft/web"
)

func GetPersonByID(ctx *web.Context) error {
	age, _ := strconv.Atoi(ctx.Request.GetParam("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Request.GetUrlParam("id_person"),
		Age:      age,
	}

	response := PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	}

	return ctx.Response.JSON(web.StatusOK, response)
}
