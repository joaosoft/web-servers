package controllers

import (
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/kataras/iris/v12"
)

func GetPersonByID(ctx iris.Context) {
	age, _ := strconv.Atoi(ctx.URLParam("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Params().Get("id_person"),
		Age:      age,
	}

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, age)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		_, _ = ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(person)
}
