package controllers

import (
	"fmt"
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

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	response := PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	}

	return ctx.JSON(http.StatusOK, response)
}
