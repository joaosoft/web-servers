package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPersonByID(ctx *gin.Context) {
	age, _ := strconv.Atoi(ctx.Request.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
		Age:      age,
	}

	ctx.Header("Content-Type", "application/json")

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.JSON(http.StatusOK, person)
}
