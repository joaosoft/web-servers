package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPersonByID(ctx *gin.Context) {
	age, _ := strconv.Atoi(ctx.Request.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param(":id_person"),
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	// ...

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	})
}
