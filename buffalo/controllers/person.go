package controllers

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"net/http"
	"strconv"
)

func GetPersonByID(ctx buffalo.Context) error {
	age, _ := strconv.Atoi(ctx.Request().URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	response := PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	}

	return ctx.Render(http.StatusOK, render.JSON(response))
}
