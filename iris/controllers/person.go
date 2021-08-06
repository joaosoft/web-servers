package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kataras/iris/v12"
)

func GetPersonByID(ctx iris.Context) {
	fmt.Println("ola")
	age, _ := strconv.Atoi(ctx.URLParam("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Params().Get("id_person"),
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	response := PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(response)
}
